import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Alert, Button, Card, Form, Container, Row, Col, Spinner, InputGroup } from 'react-bootstrap';

const CreateTransfer = () => {
  const [accounts, setAccounts] = useState([]);
  const [fromAccountID, setFromAccountID] = useState('');
  const [toAccountID, setToAccountID] = useState('');
  const [amount, setAmount] = useState('');
  const [currency, setCurrency] = useState('USD');
  const [message, setMessage] = useState('');
  const [messageType, setMessageType] = useState('info');
  const [loading, setLoading] = useState(false);
  const [transferHistory, setTransferHistory] = useState([]);

  const showMessage = (text, type = 'info') => {
    setMessage(text);
    setMessageType(type);
    setTimeout(() => setMessage(''), 5000);
  };

  const getAuthHeaders = () => {
    const token = localStorage.getItem('token');
    return token ? { Authorization: `Bearer ${token}` } : {};
  };

  const fetchAccounts = async () => {
    try {
      const response = await axios.get('/v1/accounts?page_id=1&page_size=10', {
        headers: getAuthHeaders(),
      });
      setAccounts(response.data.accounts || []);
    } catch (error) {
      console.error('Error fetching accounts:', error);
      showMessage('Please login and create accounts first', 'warning');
    }
  };

  const getSelectedAccount = (accountId) => {
    return accounts.find(acc => acc.id.toString() === accountId);
  };

  const formatBalance = (balance, currency) => {
    return new Intl.NumberFormat('en-US', {
      style: 'currency',
      currency: currency,
    }).format(balance / 100);
  };

  const validateTransfer = () => {
    if (!fromAccountID) {
      showMessage('Please select a source account', 'warning');
      return false;
    }
    if (!toAccountID) {
      showMessage('Please enter a destination account ID', 'warning');
      return false;
    }
    if (fromAccountID === toAccountID) {
      showMessage('Source and destination accounts cannot be the same', 'warning');
      return false;
    }
    if (!amount || amount <= 0) {
      showMessage('Please enter a valid amount', 'warning');
      return false;
    }
    
    const sourceAccount = getSelectedAccount(fromAccountID);
    if (!sourceAccount) {
      showMessage('Invalid source account', 'danger');
      return false;
    }
    
    if (sourceAccount.currency !== currency) {
      showMessage(`Currency mismatch. Source account uses ${sourceAccount.currency}`, 'warning');
      return false;
    }
    
    const amountInCents = amount * 100;
    if (amountInCents > sourceAccount.balance) {
      showMessage('Insufficient balance', 'danger');
      return false;
    }
    
    return true;
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    
    if (!validateTransfer()) {
      return;
    }
    
    setLoading(true);
    try {
      const response = await axios.post(
        '/v1/transfers',
        {
          from_account_id: parseInt(fromAccountID),
          to_account_id: parseInt(toAccountID),
          amount: parseInt(amount * 100), // Convert to cents
          currency: currency,
        },
        { headers: getAuthHeaders() }
      );
      
      showMessage(`Transfer successful! Transfer ID: ${response.data.transfer?.id}`, 'success');
      
      // Reset form
      setFromAccountID('');
      setToAccountID('');
      setAmount('');
      
      // Refresh accounts to show updated balances
      fetchAccounts();
      
    } catch (error) {
      console.error('Error creating transfer:', error);
      if (error.response?.status === 401) {
        showMessage('Please login first', 'warning');
      } else if (error.response?.data) {
        showMessage(`Transfer failed: ${error.response.data.message || error.response.data.error || 'Unknown error'}`, 'danger');
      } else {
        showMessage(`Transfer failed: ${error.message}`, 'danger');
      }
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchAccounts();
  }, []);

  useEffect(() => {
    if (fromAccountID) {
      const account = getSelectedAccount(fromAccountID);
      if (account) {
        setCurrency(account.currency);
      }
    }
  }, [fromAccountID, accounts]);

  const selectedAccount = getSelectedAccount(fromAccountID);

  return (
    <Container className="mt-4">
      <Row>
        <Col md={8}>
          <Card>
            <Card.Header>
              <h3>Create Transfer</h3>
            </Card.Header>
            <Card.Body>
              {message && (
                <Alert variant={messageType} dismissible onClose={() => setMessage('')}>
                  {message}
                </Alert>
              )}

              {accounts.length === 0 ? (
                <Alert variant="warning">
                  You need to create at least one account before making transfers.
                  <br />
                  <Button variant="link" href="/accounts">Go to Account Manager</Button>
                </Alert>
              ) : (
                <Form onSubmit={handleSubmit}>
                  <Form.Group className="mb-3">
                    <Form.Label>From Account</Form.Label>
                    <Form.Select
                      value={fromAccountID}
                      onChange={(e) => setFromAccountID(e.target.value)}
                      required
                    >
                      <option value="">Select source account...</option>
                      {accounts.map((account) => (
                        <option key={account.id} value={account.id}>
                          Account #{account.id} - {formatBalance(account.balance, account.currency)}
                        </option>
                      ))}
                    </Form.Select>
                    {selectedAccount && (
                      <Form.Text className="text-muted">
                        Available balance: {formatBalance(selectedAccount.balance, selectedAccount.currency)}
                      </Form.Text>
                    )}
                  </Form.Group>

                  <Form.Group className="mb-3">
                    <Form.Label>To Account ID</Form.Label>
                    <Form.Control
                      type="number"
                      placeholder="Enter destination account ID"
                      value={toAccountID}
                      onChange={(e) => setToAccountID(e.target.value)}
                      required
                    />
                    <Form.Text className="text-muted">
                      Enter the account ID of the person you want to send money to.
                    </Form.Text>
                  </Form.Group>

                  <Form.Group className="mb-3">
                    <Form.Label>Amount</Form.Label>
                    <InputGroup>
                      <InputGroup.Text>{currency}</InputGroup.Text>
                      <Form.Control
                        type="number"
                        step="0.01"
                        min="0.01"
                        placeholder="0.00"
                        value={amount}
                        onChange={(e) => setAmount(e.target.value)}
                        required
                      />
                    </InputGroup>
                    {selectedAccount && amount && (
                      <Form.Text className={amount * 100 > selectedAccount.balance ? 'text-danger' : 'text-muted'}>
                        {amount * 100 > selectedAccount.balance 
                          ? 'Insufficient balance!' 
                          : `Remaining balance after transfer: ${formatBalance(selectedAccount.balance - (amount * 100), currency)}`
                        }
                      </Form.Text>
                    )}
                  </Form.Group>

                  <Form.Group className="mb-3">
                    <Form.Label>Currency</Form.Label>
                    <Form.Control
                      type="text"
                      value={currency}
                      disabled
                      readOnly
                    />
                    <Form.Text className="text-muted">
                      Currency is automatically set based on your source account.
                    </Form.Text>
                  </Form.Group>

                  <div className="d-grid">
                    <Button type="submit" variant="primary" size="lg" disabled={loading}>
                      {loading ? (
                        <>
                          <Spinner animation="border" size="sm" className="me-2" />
                          Processing Transfer...
                        </>
                      ) : (
                        'Send Transfer'
                      )}
                    </Button>
                  </div>
                </Form>
              )}
            </Card.Body>
          </Card>
        </Col>
        
        <Col md={4}>
          <Card>
            <Card.Header>
              <h5>Transfer Tips</h5>
            </Card.Header>
            <Card.Body>
              <ul className="list-unstyled">
                <li>✓ Double-check the destination account ID</li>
                <li>✓ Ensure you have sufficient balance</li>
                <li>✓ Transfers are instant</li>
                <li>✓ Currency must match between accounts</li>
              </ul>
            </Card.Body>
          </Card>
        </Col>
      </Row>
    </Container>
  );
};

export default CreateTransfer;
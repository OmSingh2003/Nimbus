import React, { useState, useEffect } from 'react';
import { Container, Row, Col, Card, Table, Button, Alert, Form, Badge, Spinner } from 'react-bootstrap';
import { useNavigate } from 'react-router-dom';
import apiClient from '../config/api';

const Transactions = () => {
  const [transfers, setTransfers] = useState([]);
  const [accounts, setAccounts] = useState([]);
  const [selectedAccount, setSelectedAccount] = useState('');
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');
  const [currentPage, setCurrentPage] = useState(1);
  const [pageSize] = useState(10);
  const [hasMore, setHasMore] = useState(true);
  const navigate = useNavigate();

  // Fetch user accounts on component mount
  useEffect(() => {
    fetchAccounts();
  }, []);

  // Fetch transfers when account is selected
  useEffect(() => {
    if (selectedAccount) {
      fetchTransfers();
    }
  }, [selectedAccount, currentPage]);

  const fetchAccounts = async () => {
    try {
      const response = await apiClient.get('/v1/accounts', {
        params: {
          page_id: 1,
          page_size: 50
        }
      });
      
      const accounts = response.data?.accounts || [];
      setAccounts(accounts);
      if (accounts.length > 0) {
        setSelectedAccount(accounts[0].id.toString());
      }
    } catch (err) {
      setError(`Failed to fetch accounts: ${err.response?.data?.message || err.message}`);
    }
  };

  const fetchTransfers = async () => {
    if (!selectedAccount) return;
    
    setLoading(true);
    setError('');
    
    try {
      const response = await apiClient.get('/v1/transfers', {
        params: {
          page_number: currentPage,
          page_size: pageSize
        }
      });
      
      const transfers = response.data?.transfers || [];
      
      if (currentPage === 1) {
        setTransfers(transfers);
      } else {
        setTransfers(prev => [...prev, ...transfers]);
      }
      
      // Check if there are more transfers
      setHasMore(transfers && transfers.length === pageSize);
    } catch (err) {
      setError('Failed to fetch transfers');
    } finally {
      setLoading(false);
    }
  };

  const handleAccountChange = (e) => {
    setSelectedAccount(e.target.value);
    setCurrentPage(1);
    setTransfers([]);
    setHasMore(true);
  };

  const loadMore = () => {
    setCurrentPage(prev => prev + 1);
  };

  const formatAmount = (amount) => {
    return (amount / 100).toFixed(2);
  };

  const formatDate = (dateString) => {
    if (!dateString) return 'N/A';
    
    // Debug logging to understand the date format
    console.log('Date formatting debug:', { dateString, type: typeof dateString });
    
    try {
      // Handle protobuf timestamp format (for accounts)
      if (typeof dateString === 'object' && dateString.seconds) {
        const date = new Date(dateString.seconds * 1000);
        return date.toLocaleDateString() + ' ' + date.toLocaleTimeString();
      }
      
      // Handle Go time.Time string format (like "2025-07-04 13:32:33.123456 +0000 UTC")
      if (typeof dateString === 'string' && dateString.includes('UTC')) {
        // Extract the date part before the timezone
        const datePart = dateString.split(' +')[0];
        const date = new Date(datePart + 'Z'); // Add Z for UTC
        if (!isNaN(date.getTime())) {
          return date.toLocaleDateString() + ' ' + date.toLocaleTimeString();
        }
      }
      
      // Handle ISO string format
      const date = new Date(dateString);
      if (isNaN(date.getTime())) {
        return 'N/A';
      }
      
      return date.toLocaleDateString() + ' ' + date.toLocaleTimeString();
    } catch (error) {
      console.error('Date parsing error:', error, 'for dateString:', dateString);
      return 'N/A';
    }
  };

  const getTransferType = (transfer, accountId) => {
    const accountIdNum = parseInt(accountId);
    
    // Debug logging to understand the data structure
    console.log('Transfer type debug:', {
      transfer,
      accountId,
      accountIdNum,
      from_account_id: transfer.from_account_id,
      to_account_id: transfer.to_account_id
    });
    
    if (transfer.from_account_id === accountIdNum) {
      return 'outgoing';
    } else if (transfer.to_account_id === accountIdNum) {
      return 'incoming';
    }
    return 'unknown';
  };

  const getTransferBadge = (transfer, accountId) => {
    const type = getTransferType(transfer, accountId);
    if (type === 'outgoing') {
      return <Badge bg="danger">Sent</Badge>;
    } else if (type === 'incoming') {
      return <Badge bg="success">Received</Badge>;
    }
    return <Badge bg="secondary">Unknown</Badge>;
  };

  const getOtherAccountId = (transfer, accountId) => {
    const accountIdNum = parseInt(accountId);
    if (transfer.from_account_id === accountIdNum) {
      return transfer.to_account_id;
    } else if (transfer.to_account_id === accountIdNum) {
      return transfer.from_account_id;
    }
    return 'N/A';
  };

  const selectedAccountData = accounts.find(acc => acc.id.toString() === selectedAccount);

  return (
    <Container className="mt-4">
      <Row>
        <Col>
          <Card>
            <Card.Header>
              <h4 className="mb-0">Transaction History</h4>
            </Card.Header>
            <Card.Body>
              {error && (
                <Alert variant="danger" className="mb-3">
                  {error}
                </Alert>
              )}

              {/* Account Selection */}
              <Row className="mb-3">
                <Col md={6}>
                  <Form.Group>
                    <Form.Label>Select Account</Form.Label>
                    <Form.Select 
                      value={selectedAccount} 
                      onChange={handleAccountChange}
                      disabled={loading}
                    >
                      <option value="">Choose an account...</option>
                      {accounts.map(account => (
                        <option key={account.id} value={account.id.toString()}>
                          {account.currency} Account - Balance: {formatAmount(account.balance)}
                        </option>
                      ))}
                    </Form.Select>
                  </Form.Group>
                </Col>
                {selectedAccountData && (
                  <Col md={6}>
                    <div className="mt-4">
                      <h6>Current Balance</h6>
                      <h4 className="text-primary">
                        {formatAmount(selectedAccountData.balance)} {selectedAccountData.currency}
                      </h4>
                    </div>
                  </Col>
                )}
              </Row>

              {/* Transfers Table */}
              {selectedAccount && (
                <>
                  {transfers.length === 0 && !loading ? (
                    <Alert variant="info">
                      No transactions found for this account.
                    </Alert>
                  ) : (
                    <Table responsive striped hover>
                      <thead>
                        <tr>
                          <th>Date</th>
                          <th>Type</th>
                          <th>Amount</th>
                          <th>Other Account</th>
                          <th>Transfer ID</th>
                        </tr>
                      </thead>
                      <tbody>
                        {transfers.map(transfer => (
                          <tr key={transfer.id}>
                            <td>{formatDate(transfer.created_at)}</td>
                            <td>{getTransferBadge(transfer, selectedAccount)}</td>
                            <td>
                              <span className={
                                getTransferType(transfer, selectedAccount) === 'outgoing' 
                                  ? 'text-danger' 
                                  : 'text-success'
                              }>
                                {getTransferType(transfer, selectedAccount) === 'outgoing' ? '-' : '+'}
                                {formatAmount(transfer.amount)} {selectedAccountData?.currency}
                              </span>
                            </td>
                            <td>Account #{getOtherAccountId(transfer, selectedAccount)}</td>
                            <td>
                              <small className="text-muted">#{transfer.id}</small>
                            </td>
                          </tr>
                        ))}
                      </tbody>
                    </Table>
                  )}

                  {/* Load More Button */}
                  {hasMore && transfers.length > 0 && (
                    <div className="text-center mt-3">
                      <Button 
                        variant="outline-primary" 
                        onClick={loadMore}
                        disabled={loading}
                      >
                        {loading ? (
                          <>
                            <Spinner size="sm" className="me-2" />
                            Loading...
                          </>
                        ) : (
                          'Load More'
                        )}
                      </Button>
                    </div>
                  )}
                </>
              )}

              {/* Loading State */}
              {loading && transfers.length === 0 && (
                <div className="text-center py-4">
                  <Spinner animation="border" />
                  <p className="mt-2">Loading transactions...</p>
                </div>
              )}
            </Card.Body>
          </Card>
        </Col>
      </Row>
    </Container>
  );
};

export default Transactions;

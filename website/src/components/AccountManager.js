import React, { useState, useEffect } from 'react';
import apiClient, { API_CONFIG } from '../config/api';
import { Alert, Button, Card, Form, ListGroup, Container, Row, Col } from 'react-bootstrap';

const AccountManager = () => {
  const [accounts, setAccounts] = useState([]);
  const [currency, setCurrency] = useState('USD');
  const [message, setMessage] = useState('');
  const [messageType, setMessageType] = useState('info'); // success, danger, warning, info
  const [loading, setLoading] = useState(false);
  const [showCreateForm, setShowCreateForm] = useState(false);

  const showMessage = (text, type = 'info') => {
    setMessage(text);
    setMessageType(type);
    setTimeout(() => setMessage(''), 5000); // Auto-clear after 5 seconds
  };

  // Convert backend error messages to user-friendly ones
  const getUserFriendlyError = (errorMessage) => {
    if (!errorMessage) return 'An unexpected error occurred';
    
    const message = errorMessage.toLowerCase();
    
    // Authentication errors
    if (message.includes('missing authorization') || message.includes('unauthorized')) {
      return 'Please login first to manage your accounts.';
    }
    if (message.includes('invalid access token') || message.includes('token is invalid')) {
      return 'Session expired. Please login again.';
    }
    
    // Account creation errors
    if (message.includes('currency') && message.includes('invalid')) {
      return 'Please select a valid currency.';
    }
    if (message.includes('account limit') || message.includes('maximum accounts')) {
      return 'You have reached the maximum number of accounts allowed.';
    }
    
    // Database/server errors
    if (message.includes('database') || message.includes('connection')) {
      return 'Service temporarily unavailable. Please try again later.';
    }
    
    // Network errors
    if (message.includes('network') || message.includes('timeout')) {
      return 'Connection error. Please check your internet and try again.';
    }
    
    // Default fallbacks
    if (message.includes('fetch') || message.includes('load')) {
      return 'Unable to load accounts. Please refresh the page.';
    }
    if (message.includes('create')) {
      return 'Unable to create account. Please try again.';
    }
    
    return 'Something went wrong. Please try again.';
  };

  const getAuthHeaders = () => {
    const token = localStorage.getItem('token');
    const headers = token ? { Authorization: `Bearer ${token}` } : {};
    return headers;
  };

  const fetchAccounts = async () => {
    try {
      setLoading(true);
      const response = await apiClient.get('/v1/accounts?page_id=1&page_size=10');
      setAccounts(response.data.accounts || []);
      showMessage('Accounts loaded successfully', 'success');
    } catch (error) {
      let friendlyMessage;
      
      if (error.response?.status === 401) {
        friendlyMessage = 'Please login first to view your accounts.';
      } else if (error.response?.data?.message) {
        friendlyMessage = getUserFriendlyError(error.response.data.message);
      } else if (error.message) {
        friendlyMessage = getUserFriendlyError(error.message);
      } else {
        friendlyMessage = 'Unable to load accounts. Please try again.';
      }
      
      showMessage(friendlyMessage, 'danger');
    } finally {
      setLoading(false);
    }
  };

  const createAccount = async (e) => {
    e.preventDefault();
    try {
      setLoading(true);
      const response = await apiClient.post('/v1/accounts', { currency });
      
      showMessage(`Account created successfully! Account ID: ${response.data.id}`, 'success');
      setShowCreateForm(false);
      setCurrency('USD');
      fetchAccounts(); // Refresh the accounts list
    } catch (error) {
      let friendlyMessage;
      
      if (error.response?.status === 401) {
        friendlyMessage = 'Please login first to create an account.';
      } else if (error.response?.data?.message) {
        friendlyMessage = getUserFriendlyError(error.response.data.message);
      } else if (error.message) {
        friendlyMessage = getUserFriendlyError(error.message);
      } else {
        friendlyMessage = 'Unable to create account. Please try again.';
      }
      
      showMessage(friendlyMessage, 'danger');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchAccounts();
  }, []);

  const formatBalance = (balance, currency) => {
    const locale = currency === 'INR' ? 'en-IN' : 'en-US';
    return new Intl.NumberFormat(locale, {
      style: 'currency',
      currency: currency,
    }).format(balance / 100); // Assuming balance is in cents
  };

  const formatDate = (dateString) => {
    if (!dateString) return 'N/A';
    
    try {
      // Handle protobuf timestamp format (for accounts)
      if (typeof dateString === 'object' && dateString.seconds) {
        const date = new Date(dateString.seconds * 1000);
        return date.toLocaleDateString() + ' ' + date.toLocaleTimeString();
      }
      
      // Handle ISO string format
      const date = new Date(dateString);
      if (isNaN(date.getTime())) {
        return 'Invalid Date';
      }
      
      return date.toLocaleDateString() + ' ' + date.toLocaleTimeString();
    } catch (error) {
      console.error('Date parsing error:', error, 'for dateString:', dateString);
      return 'Invalid Date';
    }
  };

  return (
    <Container className="mt-4">
      <Row>
        <Col>
          <h2>Account Manager</h2>
          
          {message && (
            <Alert variant={messageType} dismissible onClose={() => setMessage('')}>
              {message}
            </Alert>
          )}

          <div className="mb-3">
            <Button 
              variant="primary" 
              onClick={() => setShowCreateForm(!showCreateForm)}
              disabled={loading}
            >
              {showCreateForm ? 'Cancel' : 'Create New Account'}
            </Button>
            <Button 
              variant="outline-secondary" 
              className="ms-2"
              onClick={fetchAccounts}
              disabled={loading}
            >
              {loading ? 'Loading...' : 'Refresh'}
            </Button>
          </div>

          {showCreateForm && (
            <Card className="mb-4">
              <Card.Header>Create New Account</Card.Header>
              <Card.Body>
                <Form onSubmit={createAccount}>
                  <Form.Group className="mb-3">
                    <Form.Label>Currency</Form.Label>
                    <Form.Select
                      value={currency}
                      onChange={(e) => setCurrency(e.target.value)}
                      required
                    >
                      <option value="USD">USD - US Dollar</option>
                      <option value="EUR">EUR - Euro</option>
                      <option value="GBP">GBP - British Pound</option>
                      <option value="CAD">CAD - Canadian Dollar</option>
                      <option value="JPY">JPY - Japanese Yen</option>
                      <option value="INR">INR - Indian Rupee</option>
                    </Form.Select>
                  </Form.Group>
                  <Button type="submit" variant="success" disabled={loading}>
                    {loading ? 'Creating...' : 'Create Account'}
                  </Button>
                </Form>
              </Card.Body>
            </Card>
          )}

          <Card>
            <Card.Header>Your Accounts</Card.Header>
            <Card.Body>
              {loading && accounts.length === 0 ? (
                <p>Loading accounts...</p>
              ) : accounts.length === 0 ? (
                <p>No accounts found. Create your first account to get started!</p>
              ) : (
                <ListGroup>
                  {accounts.map((account) => (
                    <ListGroup.Item key={account.id} className="d-flex justify-content-between align-items-start">
                      <div className="ms-2 me-auto">
                        <div className="fw-bold">Account #{account.id}</div>
                        <small className="text-muted">
                          Created: {formatDate(account.created_at)}
                        </small>
                      </div>
                      <div className="text-end">
                        <div className="fw-bold text-success">
                          {formatBalance(account.balance, account.currency)}
                        </div>
                        <small className="text-muted">{account.currency}</small>
                      </div>
                    </ListGroup.Item>
                  ))}
                </ListGroup>
              )}
            </Card.Body>
          </Card>
        </Col>
      </Row>
    </Container>
  );
};

export default AccountManager;

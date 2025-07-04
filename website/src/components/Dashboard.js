import React, { useState, useEffect } from 'react';
import { Container, Row, Col, Card, Button, Alert, Badge, Table, Spinner } from 'react-bootstrap';
import { Link, useNavigate } from 'react-router-dom';
import apiClient from '../config/api';

const Dashboard = () => {
  const [accounts, setAccounts] = useState([]);
  const [recentTransfers, setRecentTransfers] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');
  const [username, setUsername] = useState('');
  const [totalBalance, setTotalBalance] = useState({});
  const navigate = useNavigate();

  useEffect(() => {
    // Check if user is logged in
    const token = localStorage.getItem('token');
    const storedUsername = localStorage.getItem('username');
    
    if (!token) {
      // Redirect to login if no token
      navigate('/login');
      return;
    }
    
    if (storedUsername) {
      setUsername(storedUsername);
    }
    fetchDashboardData();
  }, [navigate]);

  const fetchDashboardData = async () => {
    try {
      setLoading(true);
      
      // Fetch accounts
      const accountsResponse = await apiClient.get('/accounts');
      setAccounts(accountsResponse.data || []);

      // Calculate total balance by currency
      const balances = {};
      accountsResponse.data?.forEach(account => {
        if (balances[account.currency]) {
          balances[account.currency] += account.balance;
        } else {
          balances[account.currency] = account.balance;
        }
      });
      setTotalBalance(balances);

      // Fetch recent transfers for the first account (if exists)
      if (accountsResponse.data && accountsResponse.data.length > 0) {
        try {
          const transfersResponse = await apiClient.get('/transfers', {
            params: {
              account_id: accountsResponse.data[0].id,
              page_id: 1,
              page_size: 5
            }
          });
          setRecentTransfers(transfersResponse.data || []);
        } catch (transferError) {
          console.log('No transfers found or error fetching transfers');
          setRecentTransfers([]);
        }
      }
    } catch (err) {
      setError('Failed to fetch dashboard data');
    } finally {
      setLoading(false);
    }
  };

  const formatAmount = (amount) => {
    return (amount / 100).toFixed(2);
  };

  const formatDate = (dateString) => {
    return new Date(dateString).toLocaleDateString();
  };

  const getTransferType = (transfer, accountId) => {
    if (transfer.from_account_id === accountId) {
      return 'sent';
    } else if (transfer.to_account_id === accountId) {
      return 'received';
    }
    return 'unknown';
  };

  if (loading) {
    return (
      <Container className="mt-4">
        <div className="text-center py-5">
          <Spinner animation="border" />
          <p className="mt-2">Loading your dashboard...</p>
        </div>
      </Container>
    );
  }

  return (
    <Container className="mt-4">
      {/* Welcome Header */}
      <Row className="mb-4">
        <Col>
          <div className="d-flex justify-content-between align-items-center">
            <div>
              <h2>Welcome back, {username}! 👋</h2>
              <p className="text-muted">Here's an overview of your financial activity</p>
            </div>
            <div>
              <Badge bg="success" className="p-2">
                <i className="bi bi-shield-check"></i> Account Verified
              </Badge>
            </div>
          </div>
        </Col>
      </Row>

      {error && (
        <Alert variant="danger" className="mb-4">
          {error}
        </Alert>
      )}

      {/* Total Balance Cards */}
      <Row className="g-4 mb-4">
        <Col md={8}>
          <Card className="h-100">
            <Card.Header className="bg-primary text-white">
              <h5 className="mb-0">💰 Total Balance Overview</h5>
            </Card.Header>
            <Card.Body>
              {Object.keys(totalBalance).length > 0 ? (
                <Row>
                  {Object.entries(totalBalance).map(([currency, balance]) => (
                    <Col md={6} key={currency} className="mb-3">
                      <div className="text-center p-3 bg-light rounded">
                        <h4 className="text-primary mb-1">
                          {formatAmount(balance)} {currency}
                        </h4>
                        <small className="text-muted">Total in {currency}</small>
                      </div>
                    </Col>
                  ))}
                </Row>
              ) : (
                <div className="text-center py-4">
                  <p className="text-muted">No accounts created yet</p>
                  <Button as={Link} to="/accounts" variant="primary">
                    Create Your First Account
                  </Button>
                </div>
              )}
            </Card.Body>
          </Card>
        </Col>
        <Col md={4}>
          <Card className="h-100 bg-gradient-primary text-white">
            <Card.Body className="d-flex flex-column justify-content-center text-center">
              <h3 className="mb-2">{accounts.length}</h3>
              <p className="mb-1">Active Accounts</p>
              <small>Across multiple currencies</small>
            </Card.Body>
          </Card>
        </Col>
      </Row>

      {/* Quick Actions */}
      <Row className="g-4 mb-4">
        <Col>
          <Card>
            <Card.Header>
              <h5 className="mb-0">🚀 Quick Actions</h5>
            </Card.Header>
            <Card.Body>
              <Row className="g-3">
                <Col md={3}>
                  <div className="d-grid">
                    <Button as={Link} to="/accounts" variant="outline-primary" size="lg">
                      <div className="py-2">
                        <div>💼</div>
                        <div>Manage Accounts</div>
                      </div>
                    </Button>
                  </div>
                </Col>
                <Col md={3}>
                  <div className="d-grid">
                    <Button as={Link} to="/transfer" variant="outline-success" size="lg">
                      <div className="py-2">
                        <div>💸</div>
                        <div>Send Money</div>
                      </div>
                    </Button>
                  </div>
                </Col>
                <Col md={3}>
                  <div className="d-grid">
                    <Button as={Link} to="/transactions" variant="outline-info" size="lg">
                      <div className="py-2">
                        <div>📊</div>
                        <div>View Transactions</div>
                      </div>
                    </Button>
                  </div>
                </Col>
                <Col md={3}>
                  <div className="d-grid">
                    <Button as={Link} to="/docs" variant="outline-secondary" size="lg">
                      <div className="py-2">
                        <div>📖</div>
                        <div>Documentation</div>
                      </div>
                    </Button>
                  </div>
                </Col>
              </Row>
            </Card.Body>
          </Card>
        </Col>
      </Row>

      {/* Accounts Overview */}
      <Row className="g-4 mb-4">
        <Col md={8}>
          <Card>
            <Card.Header className="d-flex justify-content-between align-items-center">
              <h5 className="mb-0">🏦 Your Accounts</h5>
              <Button as={Link} to="/accounts" variant="outline-primary" size="sm">
                Manage All
              </Button>
            </Card.Header>
            <Card.Body>
              {accounts.length > 0 ? (
                <Table responsive>
                  <thead>
                    <tr>
                      <th>Currency</th>
                      <th>Balance</th>
                      <th>Status</th>
                    </tr>
                  </thead>
                  <tbody>
                    {accounts.slice(0, 4).map(account => (
                      <tr key={account.id}>
                        <td>
                          <strong>{account.currency}</strong>
                        </td>
                        <td>
                          <span className="text-success">
                            {formatAmount(account.balance)} {account.currency}
                          </span>
                        </td>
                        <td>
                          <Badge bg="success">Active</Badge>
                        </td>
                      </tr>
                    ))}
                  </tbody>
                </Table>
              ) : (
                <div className="text-center py-4">
                  <p className="text-muted">No accounts found</p>
                  <Button as={Link} to="/accounts" variant="primary">
                    Create Account
                  </Button>
                </div>
              )}
            </Card.Body>
          </Card>
        </Col>
        
        {/* Recent Activity */}
        <Col md={4}>
          <Card>
            <Card.Header className="d-flex justify-content-between align-items-center">
              <h5 className="mb-0">📈 Recent Activity</h5>
              <Button as={Link} to="/transactions" variant="outline-primary" size="sm">
                View All
              </Button>
            </Card.Header>
            <Card.Body>
              {recentTransfers.length > 0 ? (
                <div>
                  {recentTransfers.slice(0, 3).map(transfer => (
                    <div key={transfer.id} className="d-flex justify-content-between align-items-center mb-3 pb-2 border-bottom">
                      <div>
                        <div className="small text-muted">
                          {formatDate(transfer.created_at)}
                        </div>
                        <div>
                          {getTransferType(transfer, accounts[0]?.id) === 'sent' ? (
                            <Badge bg="danger" className="me-1">Sent</Badge>
                          ) : (
                            <Badge bg="success" className="me-1">Received</Badge>
                          )}
                          <small>Transfer #{transfer.id}</small>
                        </div>
                      </div>
                      <div className="text-end">
                        <div className={
                          getTransferType(transfer, accounts[0]?.id) === 'sent' 
                            ? 'text-danger' 
                            : 'text-success'
                        }>
                          {getTransferType(transfer, accounts[0]?.id) === 'sent' ? '-' : '+'}
                          {formatAmount(transfer.amount)}
                        </div>
                      </div>
                    </div>
                  ))}
                </div>
              ) : (
                <div className="text-center py-3">
                  <p className="text-muted small">No recent transactions</p>
                  <Button as={Link} to="/transfer" variant="outline-success" size="sm">
                    Make Transfer
                  </Button>
                </div>
              )}
            </Card.Body>
          </Card>
        </Col>
      </Row>

      {/* Tips & Security */}
      <Row className="g-4">
        <Col md={6}>
          <Card className="border-info">
            <Card.Header className="bg-info text-white">
              <h6 className="mb-0">💡 Pro Tips</h6>
            </Card.Header>
            <Card.Body>
              <ul className="mb-0">
                <li className="mb-2">Create accounts in multiple currencies for better exchange rates</li>
                <li className="mb-2">Use instant transfers for quick payments between your accounts</li>
                <li className="mb-0">Check your transaction history regularly for security</li>
              </ul>
            </Card.Body>
          </Card>
        </Col>
        <Col md={6}>
          <Card className="border-warning">
            <Card.Header className="bg-warning text-dark">
              <h6 className="mb-0">🔒 Security Status</h6>
            </Card.Header>
            <Card.Body>
              <div className="d-flex align-items-center mb-2">
                <Badge bg="success" className="me-2">✓</Badge>
                <span>Account verified</span>
              </div>
              <div className="d-flex align-items-center mb-2">
                <Badge bg="success" className="me-2">✓</Badge>
                <span>JWT token authentication</span>
              </div>
              <div className="d-flex align-items-center">
                <Badge bg="success" className="me-2">✓</Badge>
                <span>Encrypted data transmission</span>
              </div>
            </Card.Body>
          </Card>
        </Col>
      </Row>
    </Container>
  );
};

export default Dashboard;

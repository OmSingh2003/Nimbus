import React, { useState, useEffect } from 'react';
import { Container, Row, Col, Card, Button, Alert } from 'react-bootstrap';
import { Link } from 'react-router-dom';

const Home = () => {
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [username, setUsername] = useState('');

  useEffect(() => {
    const token = localStorage.getItem('token');
    const storedUsername = localStorage.getItem('username');
    if (token) {
      setIsLoggedIn(true);
      setUsername(storedUsername || 'User');
    }
  }, []);

  return (
    <Container className="mt-4">
      {/* Hero Section */}
      <Row>
        <Col>
          <div className="hero-section text-center mb-5">
            <h1 className="display-3 mb-3">Invest in everything</h1>
            <p className="lead fs-4">Online platform to invest in stocks, derivatives, mutual funds, ETFs, bonds, and more.</p>
            <p className="text-muted">Experience modern banking with enterprise-grade security and zero-fee account opening</p>
            
            {isLoggedIn ? (
              <Alert variant="success" className="mt-4">
                Welcome back, <strong>{username}</strong>! 
                You're all set to manage your accounts.
              </Alert>
            ) : (
              <div className="mt-4">
                <Button as={Link} to="/create-user" variant="primary" size="lg" className="me-3">
                  Open account
                </Button>
                <Button as={Link} to="/login" variant="outline-primary" size="lg">
                  Sign in
                </Button>
              </div>
            )}
          </div>
        </Col>
      </Row>

      {/* Stats Section */}
      <Row className="mb-5">
        <Col>
          <div className="stats-section">
            <Container>
              <Row>
                <Col md={3}>
                  <div className="stat-item">
                    <span className="stat-number">1M+</span>
                    <div className="stat-label">Active clients</div>
                  </div>
                </Col>
                <Col md={3}>
                  <div className="stat-item">
                    <span className="stat-number">₹6+ lakh cr</span>
                    <div className="stat-label">Client assets</div>
                  </div>
                </Col>
                <Col md={3}>
                  <div className="stat-item">
                    <span className="stat-number">Zero</span>
                    <div className="stat-label">Account opening charges</div>
                  </div>
                </Col>
                <Col md={3}>
                  <div className="stat-item">
                    <span className="stat-number">₹20</span>
                    <div className="stat-label">Maximum brokerage</div>
                  </div>
                </Col>
              </Row>
            </Container>
          </div>
        </Col>
      </Row>

      {/* Quick Actions for Logged In Users */}
      {isLoggedIn && (
        <Row className="g-4 mb-5">
          <Col md={4}>
            <Card className="h-100 border-primary">
              <Card.Body className="text-center">
                <div className="feature-icon">💼</div>
                <Card.Title>My Accounts</Card.Title>
                <Card.Text>
                  View and manage your bank accounts, check balances, and create new accounts.
                </Card.Text>
                <Button as={Link} to="/accounts" variant="primary">
                  Manage Accounts
                </Button>
              </Card.Body>
            </Card>
          </Col>
          <Col md={4}>
            <Card className="h-100 border-success">
              <Card.Body className="text-center">
                <div className="feature-icon">💸</div>
                <Card.Title>Send Money</Card.Title>
                <Card.Text>
                  Transfer money securely between accounts with our instant transfer system.
                </Card.Text>
                <Button as={Link} to="/transfer" variant="success">
                  Make Transfer
                </Button>
              </Card.Body>
            </Card>
          </Col>
          <Col md={4}>
            <Card className="h-100 border-info">
              <Card.Body className="text-center">
                <div className="feature-icon">🔐</div>
                <Card.Title>Account Security</Card.Title>
                <Card.Text>
                  Your security is our priority. All transactions are encrypted and monitored.
                </Card.Text>
                <Button variant="info" disabled>
                  Protected ✓
                </Button>
              </Card.Body>
            </Card>
          </Col>
        </Row>
      )}

      {/* Features Section */}
      <Row className="mb-5">
        <Col>
          <div className="features-section">
            <h2 className="text-center mb-4">Why choose VaultGuard?</h2>
            <Row className="g-4">
              <Col md={4}>
                <div className="feature-card">
                  <div className="feature-icon">🏦</div>
                  <h5>Multi-Currency Support</h5>
                  <p>Create accounts in USD, EUR, GBP, CAD, JPY, and INR. Manage global finances from one platform.</p>
                </div>
              </Col>
              <Col md={4}>
                <div className="feature-card">
                  <div className="feature-icon">⚡</div>
                  <h5>Instant Transfers</h5>
                  <p>Send money instantly between accounts with real-time balance updates and zero processing delays.</p>
                </div>
              </Col>
              <Col md={4}>
                <div className="feature-card">
                  <div className="feature-icon">🔒</div>
                  <h5>Bank-Level Security</h5>
                  <p>Your data is encrypted and protected with industry-standard security protocols and 2FA.</p>
                </div>
              </Col>
              <Col md={4}>
                <div className="feature-card">
                  <div className="feature-icon">📱</div>
                  <h5>Mobile Ready</h5>
                  <p>Access your accounts anywhere, anytime with our responsive web platform optimized for all devices.</p>
                </div>
              </Col>
              <Col md={4}>
                <div className="feature-card">
                  <div className="feature-icon">💹</div>
                  <h5>Real-time Analytics</h5>
                  <p>Track your spending patterns, account balances, and transaction history with detailed insights.</p>
                </div>
              </Col>
              <Col md={4}>
                <div className="feature-card">
                  <div className="feature-icon">🌍</div>
                  <h5>Global Reach</h5>
                  <p>Bank globally with support for international transfers and multi-currency account management.</p>
                </div>
              </Col>
            </Row>
          </div>
        </Col>
      </Row>

      {/* Technology Section */}
      <Row className="mb-5">
        <Col>
          <Card>
            <Card.Body className="text-center p-5">
              <h2 className="mb-4">Built with modern technology</h2>
              <Row>
                <Col md={3}>
                  <h5>🚀 Go Backend</h5>
                  <p>High-performance server built with Go for speed and reliability</p>
                </Col>
                <Col md={3}>
                  <h5>⚛️ React Frontend</h5>
                  <p>Modern, responsive user interface built with React and Bootstrap</p>
                </Col>
                <Col md={3}>
                  <h5>🗄️ PostgreSQL</h5>
                  <p>Robust database solution for secure and scalable data management</p>
                </Col>
                <Col md={3}>
                  <h5>🔐 JWT Security</h5>
                  <p>Secure authentication with JSON Web Tokens and session management</p>
                </Col>
              </Row>
            </Card.Body>
          </Card>
        </Col>
      </Row>

      {/* Getting Started Section */}
      {!isLoggedIn && (
        <Row className="mb-5">
          <Col>
            <Card className="border-primary">
              <Card.Body className="text-center p-5">
                <h2 className="mb-4">Ready to get started?</h2>
                <p className="lead mb-4">
                  Join thousands of users who trust VaultGuard for their banking needs. 
                  Open your account today and experience the future of banking.
                </p>
                <div>
                  <Button as={Link} to="/create-user" variant="primary" size="lg" className="me-3">
                    Create Account - It's Free
                  </Button>
                  <Button as={Link} to="/login" variant="outline-primary" size="lg">
                    Already have account? Login
                  </Button>
                </div>
                <div className="mt-4">
                  <small className="text-muted">
                    ✓ No account maintenance charges &nbsp;&nbsp;
                    ✓ Free debit card &nbsp;&nbsp;
                    ✓ 24/7 customer support
                  </small>
                </div>
              </Card.Body>
            </Card>
          </Col>
        </Row>
      )}
    </Container>
  );
};

export default Home;

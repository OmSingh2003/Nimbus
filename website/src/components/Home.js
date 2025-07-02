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

  const handleLogout = () => {
    localStorage.removeItem('token');
    localStorage.removeItem('username');
    setIsLoggedIn(false);
    setUsername('');
  };

  return (
    <Container className="mt-4">
      <Row>
        <Col>
          <div className="hero-section text-center mb-5">
            <h1 className="display-4 text-primary">Welcome to VaultGuard</h1>
            <p className="lead">Secure Banking Made Simple</p>
            {isLoggedIn && (
              <Alert variant="success" className="mt-3">
                Welcome back, <strong>{username}</strong>! 
                <Button variant="outline-danger" size="sm" className="ms-3" onClick={handleLogout}>
                  Logout
                </Button>
              </Alert>
            )}
          </div>
        </Col>
      </Row>

      <Row className="g-4">
        {!isLoggedIn ? (
          <>
            <Col md={6}>
              <Card className="h-100">
                <Card.Body className="text-center">
                  <Card.Title>New User?</Card.Title>
                  <Card.Text>
                    Create your secure VaultGuard account and start managing your finances today.
                  </Card.Text>
                  <Button as={Link} to="/create-user" variant="primary">
                    Create Account
                  </Button>
                </Card.Body>
              </Card>
            </Col>
            <Col md={6}>
              <Card className="h-100">
                <Card.Body className="text-center">
                  <Card.Title>Existing User?</Card.Title>
                  <Card.Text>
                    Login to access your accounts, view balances, and make transfers.
                  </Card.Text>
                  <Button as={Link} to="/login" variant="success">
                    Login
                  </Button>
                </Card.Body>
              </Card>
            </Col>
          </>
        ) : (
          <>
            <Col md={4}>
              <Card className="h-100 border-primary">
                <Card.Body className="text-center">
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
          </>
        )}
      </Row>

      <Row className="mt-5">
        <Col>
          <Card>
            <Card.Header>
              <h4>Features</h4>
            </Card.Header>
            <Card.Body>
              <Row>
                <Col md={4}>
                  <h5>🏦 Multi-Currency Support</h5>
                  <p>Create accounts in USD, EUR, GBP, CAD, and JPY.</p>
                </Col>
                <Col md={4}>
                  <h5>⚡ Instant Transfers</h5>
                  <p>Send money instantly between accounts with real-time balance updates.</p>
                </Col>
                <Col md={4}>
                  <h5>🔒 Bank-Level Security</h5>
                  <p>Your data is encrypted and protected with industry-standard security.</p>
                </Col>
              </Row>
            </Card.Body>
          </Card>
        </Col>
      </Row>
    </Container>
  );
};

export default Home;

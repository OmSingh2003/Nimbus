import React, { useState } from 'react';
import axios from 'axios';
import { Alert, Button, Card, Form, Container, Row, Col, Spinner } from 'react-bootstrap';
import { useNavigate } from 'react-router-dom';

const LoginUser = ({ onLogin }) => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [message, setMessage] = useState('');
  const [messageType, setMessageType] = useState('info');
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();

  const showMessage = (text, type = 'info') => {
    setMessage(text);
    setMessageType(type);
    setTimeout(() => setMessage(''), 5000);
  };

  // Convert backend error messages to user-friendly ones
  const getUserFriendlyError = (errorMessage) => {
    if (!errorMessage) return 'Login failed. Please try again.';
    
    const message = errorMessage.toLowerCase();
    
    // Authentication errors
    if (message.includes('user not found') || message.includes('no rows in result set')) {
      return 'Username not found. Please check your username or create an account.';
    }
    if (message.includes('invalid password') || message.includes('incorrect password')) {
      return 'Incorrect password. Please try again.';
    }
    if (message.includes('invalid credentials')) {
      return 'Invalid username or password. Please try again.';
    }
    
    // Token/session errors
    if (message.includes('token') && message.includes('invalid')) {
      return 'Session expired. Please login again.';
    }
    
    // Account status errors
    if (message.includes('account') && message.includes('disabled')) {
      return 'Account is disabled. Please contact support.';
    }
    if (message.includes('account') && message.includes('locked')) {
      return 'Account is temporarily locked. Please try again later.';
    }
    
    // Network/connection errors
    if (message.includes('network') || message.includes('connection')) {
      return 'Connection error. Please check your internet and try again.';
    }
    
    // Default fallback
    return 'Login failed. Please check your credentials and try again.';
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);
    
    try {
      const response = await axios.post('/v1/login_user', {
        username,
        password,
      });
      
      localStorage.setItem('token', response.data.accessToken);
      localStorage.setItem('username', username);
      
      // Call the onLogin prop to update App state
      if (onLogin) {
        onLogin(username);
      }
      
      showMessage(`Welcome back, ${username}!`, 'success');
      
      // Navigate to accounts page after successful login
      setTimeout(() => {
        navigate('/accounts');
      }, 1000);
      
    } catch (error) {
      console.error('Error logging in:', error);
      let friendlyMessage;
      
      if (error.response?.data?.message) {
        friendlyMessage = getUserFriendlyError(error.response.data.message);
      } else if (error.response?.data?.error) {
        friendlyMessage = getUserFriendlyError(error.response.data.error);
      } else if (error.message) {
        friendlyMessage = getUserFriendlyError(error.message);
      } else {
        friendlyMessage = 'Login failed. Please try again.';
      }
      
      showMessage(friendlyMessage, 'danger');
    } finally {
      setLoading(false);
    }
  };

  return (
    <Container className="mt-4">
      <Row className="justify-content-center">
        <Col md={6}>
          <Card>
            <Card.Header>
              <h3 className="text-center">Login to VaultGuard</h3>
            </Card.Header>
            <Card.Body>
              {message && (
                <Alert variant={messageType} dismissible onClose={() => setMessage('')}>
                  {message}
                </Alert>
              )}
              
              <Form onSubmit={handleSubmit}>
                <Form.Group className="mb-3">
                  <Form.Label>Username</Form.Label>
                  <Form.Control
                    type="text"
                    placeholder="Enter username"
                    value={username}
                    onChange={(e) => setUsername(e.target.value)}
                    required
                  />
                </Form.Group>

                <Form.Group className="mb-3">
                  <Form.Label>Password</Form.Label>
                  <Form.Control
                    type="password"
                    placeholder="Enter password"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                    required
                  />
                </Form.Group>

                <div className="d-grid">
                  <Button type="submit" variant="primary" size="lg" disabled={loading}>
                    {loading ? (
                      <>
                        <Spinner animation="border" size="sm" className="me-2" />
                        Logging in...
                      </>
                    ) : (
                      'Login'
                    )}
                  </Button>
                </div>
              </Form>
              
              <div className="text-center mt-3">
                <small className="text-muted">
                  Don't have an account? <a href="/create-user">Create one here</a>
                </small>
              </div>
            </Card.Body>
          </Card>
        </Col>
      </Row>
    </Container>
  );
};

export default LoginUser;
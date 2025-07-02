import React, { useState } from 'react';
import apiClient, { API_CONFIG } from '../config/api';
import { Alert, Button, Card, Form, Container, Row, Col, Spinner } from 'react-bootstrap';

const CreateUser = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [fullName, setFullName] = useState('');
  const [email, setEmail] = useState('');
  const [message, setMessage] = useState('');
  const [messageType, setMessageType] = useState('info');
  const [loading, setLoading] = useState(false);

  const showMessage = (text, type = 'info') => {
    setMessage(text);
    setMessageType(type);
    setTimeout(() => setMessage(''), 7000);
  };

  // Convert backend error messages to user-friendly ones
  const getUserFriendlyError = (errorMessage) => {
    if (!errorMessage) return 'An unexpected error occurred';
    
    const message = errorMessage.toLowerCase();
    
    // Username/duplicate errors
    if (message.includes('duplicate key') && message.includes('users_pkey')) {
      return 'Username already exists. Please choose a different username.';
    }
    if (message.includes('duplicate key') && message.includes('email')) {
      return 'Email address already registered. Please use a different email.';
    }
    
    // Password validation errors
    if (message.includes('password') && (message.includes('length') || message.includes('short'))) {
      return 'Password must be at least 8 characters long.';
    }
    if (message.includes('password') && message.includes('uppercase')) {
      return 'Password must include at least one uppercase letter.';
    }
    if (message.includes('password') && message.includes('lowercase')) {
      return 'Password must include at least one lowercase letter.';
    }
    if (message.includes('password') && message.includes('number')) {
      return 'Password must include at least one number.';
    }
    if (message.includes('password') && message.includes('special')) {
      return 'Password must include at least one special character (!@#$%^&*).';
    }
    
    // Email validation errors
    if (message.includes('email') && message.includes('invalid')) {
      return 'Please enter a valid email address.';
    }
    
    // Username validation errors
    if (message.includes('username') && message.includes('invalid')) {
      return 'Username can only contain letters, numbers, and underscores.';
    }
    if (message.includes('username') && (message.includes('short') || message.includes('length'))) {
      return 'Username must be at least 3 characters long.';
    }
    
    // Network/connection errors
    if (message.includes('network') || message.includes('connection')) {
      return 'Connection error. Please check your internet and try again.';
    }
    
    // Default fallback
    return 'Unable to create account. Please check your information and try again.';
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);
    
    try {
      const response = await apiClient.post(API_CONFIG.ENDPOINTS.CREATE_USER, {
        username,
        password,
        full_name: fullName,
        email,
      });
      
      showMessage(`Account created successfully! Please check your email to verify your account before logging in. A verification link has been sent to ${email}.`, 'success');
      
      // Clear form
      setUsername('');
      setPassword('');
      setFullName('');
      setEmail('');
      
    } catch (error) {
      console.error('Error creating user:', error);
      let friendlyMessage;
      
      if (error.response?.data?.message) {
        friendlyMessage = getUserFriendlyError(error.response.data.message);
      } else if (error.response?.data?.error) {
        friendlyMessage = getUserFriendlyError(error.response.data.error);
      } else if (error.message) {
        friendlyMessage = getUserFriendlyError(error.message);
      } else {
        friendlyMessage = 'An unexpected error occurred. Please try again.';
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
              <h3 className="text-center">Create New Account</h3>
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
                  <Form.Text className="text-muted">
                    Choose a unique username (letters, numbers, and underscores only).
                  </Form.Text>
                </Form.Group>

                <Form.Group className="mb-3">
                  <Form.Label>Full Name</Form.Label>
                  <Form.Control
                    type="text"
                    placeholder="Enter your full name"
                    value={fullName}
                    onChange={(e) => setFullName(e.target.value)}
                    required
                  />
                </Form.Group>

                <Form.Group className="mb-3">
                  <Form.Label>Email</Form.Label>
                  <Form.Control
                    type="email"
                    placeholder="Enter email address"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                    required
                  />
                  <Form.Text className="text-muted">
                    We'll send account verification to this email.
                  </Form.Text>
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
                  <Form.Text className="text-muted">
                    Password must include uppercase, lowercase, number, and special character.
                  </Form.Text>
                </Form.Group>

                <div className="d-grid">
                  <Button type="submit" variant="primary" size="lg" disabled={loading}>
                    {loading ? (
                      <>
                        <Spinner animation="border" size="sm" className="me-2" />
                        Creating Account...
                      </>
                    ) : (
                      'Create Account'
                    )}
                  </Button>
                </div>
              </Form>
            </Card.Body>
          </Card>
        </Col>
      </Row>
    </Container>
  );
};

export default CreateUser;
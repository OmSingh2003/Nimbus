import React, { useState, useEffect } from 'react';
import { useSearchParams, useNavigate } from 'react-router-dom';
import { Container, Row, Col, Card, Alert, Button, Spinner } from 'react-bootstrap';
import axios from 'axios';

const VerifyEmail = () => {
  const [searchParams] = useSearchParams();
  const navigate = useNavigate();
  const [status, setStatus] = useState('verifying'); // 'verifying', 'success', 'error'
  const [message, setMessage] = useState('');
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const emailId = searchParams.get('email_id');
    const secretCode = searchParams.get('secret_code');

    if (!emailId || !secretCode) {
      setStatus('error');
      setMessage('Invalid verification link. Missing required parameters.');
      setLoading(false);
      return;
    }

    const verifyEmail = async () => {
      try {
        const response = await axios.post('/v1/verify_email', {
          email_id: parseInt(emailId),
          secret_code: secretCode,
        });

        setStatus('success');
        setMessage(response.data.message || 'Email verified successfully! Welcome to Nimbus! A $100 USD welcome account has been created for you.');
        setLoading(false);
      } catch (error) {
        console.error('Email verification failed:', error);
        let errorMessage = 'Email verification failed. ';
        
        if (error.response?.data?.message) {
          errorMessage += error.response.data.message;
        } else if (error.response?.status === 400) {
          errorMessage += 'Invalid or expired verification link.';
        } else {
          errorMessage += 'Please try again or contact support.';
        }
        
        setStatus('error');
        setMessage(errorMessage);
        setLoading(false);
      }
    };

    verifyEmail();
  }, [searchParams]);

  const handleContinue = () => {
    if (status === 'success') {
      navigate('/login');
    } else {
      navigate('/');
    }
  };

  return (
    <Container className="mt-5">
      <Row className="justify-content-center">
        <Col md={8} lg={6}>
          <Card>
            <Card.Header className="text-center">
              <h3>Email Verification</h3>
            </Card.Header>
            <Card.Body className="text-center p-5">
              {loading ? (
                <div>
                  <Spinner animation="border" variant="primary" className="mb-3" />
                  <p>Verifying your email address...</p>
                </div>
              ) : (
                <div>
                  {status === 'success' && (
                    <div>
                      <div className="mb-4">
                        <svg 
                          width="64" 
                          height="64" 
                          fill="#00c851" 
                          viewBox="0 0 16 16" 
                          className="mb-3"
                        >
                          <path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14zm0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16z"/>
                          <path d="M10.97 4.97a.235.235 0 0 0-.02.022L7.477 9.417 5.384 7.323a.75.75 0 0 0-1.06 1.06L6.97 11.03a.75.75 0 0 0 1.079-.02l3.992-4.99a.75.75 0 0 0-1.071-1.05z"/>
                        </svg>
                      </div>
                      <Alert variant="success">
                        <Alert.Heading>🎉 Welcome to Nimbus!</Alert.Heading>
                        <p>{message}</p>
                        <hr />
                        <p className="mb-0">
                          <strong>What's next?</strong><br/>
                          • Log in to your account<br/>
                          • Explore your $100 USD welcome account<br/>
                          • Create additional accounts in different currencies<br/>
                          • Start making secure transactions
                        </p>
                      </Alert>
                    </div>
                  )}
                  
                  {status === 'error' && (
                    <div>
                      <div className="mb-4">
                        <svg 
                          width="64" 
                          height="64" 
                          fill="#e74c3c" 
                          viewBox="0 0 16 16" 
                          className="mb-3"
                        >
                          <path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14zm0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16z"/>
                          <path d="M4.646 4.646a.5.5 0 0 1 .708 0L8 7.293l2.646-2.647a.5.5 0 0 1 .708.708L8.707 8l2.647 2.646a.5.5 0 0 1-.708.708L8 8.707l-2.646 2.647a.5.5 0 0 1-.708-.708L7.293 8 4.646 5.354a.5.5 0 0 1 0-.708z"/>
                        </svg>
                      </div>
                      <Alert variant="danger">
                        <Alert.Heading>Verification Failed</Alert.Heading>
                        <p>{message}</p>
                        <hr />
                        <p className="mb-0">
                          <strong>What can you do?</strong><br/>
                          • Try creating a new account<br/>
                          • Contact support if you continue having issues<br/>
                          • Check if the link has expired (links expire after 24 hours)
                        </p>
                      </Alert>
                    </div>
                  )}

                  <div className="mt-4">
                    <Button 
                      variant={status === 'success' ? 'primary' : 'secondary'}
                      size="lg"
                      onClick={handleContinue}
                    >
                      {status === 'success' ? 'Continue to Login' : 'Back to Home'}
                    </Button>
                  </div>
                </div>
              )}
            </Card.Body>
          </Card>
        </Col>
      </Row>
    </Container>
  );
};

export default VerifyEmail;

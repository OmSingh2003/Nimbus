import React, { useState } from 'react';
import { Container, Row, Col, Card, Badge, Button, Alert, Tab, Tabs, Table } from 'react-bootstrap';

const Documentation = () => {
  const [demoAccountNumber, setDemoAccountNumber] = useState('DEMO-1234567890');

  return (
    <Container className="mt-4">
      <Row>
        <Col>
          <div className="text-center mb-5">
            <h1 className="display-4">Documentation</h1>
            <p className="lead">Complete guide to the Nimbus secure banking platform</p>
            <Badge bg="primary" className="me-2">v1.0.0</Badge>
            <Badge bg="success" className="me-2">Production Ready</Badge>
            <Badge bg="info">Go + React + PostgreSQL</Badge>
          </div>
        </Col>
      </Row>

      <Tabs defaultActiveKey="overview" className="mb-4">
        {/* Overview Tab */}
        <Tab eventKey="overview" title="Overview">
          <Row>
            <Col lg={8} className="mx-auto">
                  <h3><img src="/icon.png" alt="About" width="32" height="32" className="me-2" />About Nimbus</h3>
                  <p>
                    Nimbus is a modern, secure banking platform built with enterprise-grade technologies. 
                    It provides comprehensive financial services including multi-currency accounts, 
                    instant transfers, and robust security features.
                  </p>
                  
                  <h4>Key Features</h4>
                  <ul>
                    <li><strong>Multi-Currency Support:</strong> USD, EUR, GBP, CAD, JPY, INR</li>
                    <li><strong>Email Verification:</strong> Secure account activation with 24-hour verification links</li>
                    <li><strong>Welcome Credits:</strong> $100 USD equivalent in user's preferred currency</li>
                    <li><strong>Instant Transfers:</strong> Real-time money transfers between accounts</li>
                    <li><strong>JWT Authentication:</strong> Secure session management</li>
                    <li><strong>Responsive Design:</strong> Works seamlessly across devices</li>
                  </ul>

                  <h4>Target Audience</h4>
                  <p>
                    Designed for individuals and businesses seeking modern, secure digital banking 
                    with multi-currency capabilities and instant transaction processing.
                  </p>
            </Col>
          </Row>
        </Tab>

        {/* Technical Stack Tab */}
        <Tab eventKey="tech-stack" title="Technical Stack">
          <Row>
            <Col md={6}>
              <h4><img src="/golang.png" alt="Golang" width="24" height="24" className="me-2" /> Backend Technologies</h4>
              <ul>
                <li><strong>Language:</strong> Go (Golang)</li>
                <li><strong>Framework:</strong> Gin HTTP Framework</li>
                <li><strong>Database:</strong> PostgreSQL</li>
                <li><strong>ORM:</strong> SQLC (SQL Compiler)</li>
                <li><strong>Authentication:</strong> JWT Tokens</li>
                <li><strong>API:</strong> RESTful + gRPC</li>
                <li><strong>Email Service:</strong> SMTP Integration</li>
                <li><strong>Queue System:</strong> Redis + Asynq</li>
              </ul>
            </Col>
            <Col md={6}>
              <h4><img src="/javaScipt.png" alt="JavaScript" width="24" height="24" className="me-2" /> Frontend Technologies</h4>
              <ul>
                <li><strong>Framework:</strong> React 18</li>
                <li><strong>Routing:</strong> React Router v6</li>
                <li><strong>UI Library:</strong> React Bootstrap</li>
                <li><strong>HTTP Client:</strong> Axios</li>
                <li><strong>State Management:</strong> React Hooks</li>
                <li><strong>Styling:</strong> Custom CSS + Bootstrap</li>
                <li><strong>Build Tool:</strong> Create React App</li>
                <li><strong>Icons:</strong> Bootstrap Icons</li>
              </ul>
            </Col>
          </Row>

          <Row>
            <Col>
              <Card>
                <Card.Header><h4><img src="/microservices.png" alt="Microservices" width="24" height="24" className="me-2" /> Microservices Architecture</h4></Card.Header>
                <Card.Body>
                  <p><strong>Microservices Architecture:</strong></p>
                  <ul>
                    <li><strong>API Gateway:</strong> Handles routing and authentication</li>
                    <li><strong>User Service:</strong> Account management and verification</li>
                    <li><strong>Transaction Service:</strong> Money transfers and balance management</li>
                    <li><strong>Email Service:</strong> Verification and notification emails</li>
                    <li><strong>Database Layer:</strong> PostgreSQL with connection pooling</li>
                  </ul>
                </Card.Body>
              </Card>
            </Col>
          </Row>
        </Tab>

        {/* User Guide Tab */}
        <Tab eventKey="user-guide" title="User Guide">
          <Row>
            <Col lg={10} className="mx-auto">
              <Card className="mb-4">
                <Card.Header><h3>📋 Getting Started</h3></Card.Header>
                <Card.Body>
                  <div className="mb-4">
                    <h4>1. 🎯 Account Registration</h4>
                    <ol>
                      <li>Click "Create Account" in the navigation</li>
                      <li>Fill in your details (username, full name, email, password)</li>
                      <li>Submit the form and check your email</li>
                      <li>Click the verification link (expires in 24 hours)</li>
                      <li>Welcome! You'll receive $100 USD in your account</li>
                    </ol>
                  </div>

                  <div className="mb-4">
                    <h4>2. 🔐 Login Process</h4>
                    <ol>
                      <li>Use your username and password to login</li>
                      <li>Unverified accounts cannot login (check email first)</li>
                      <li>Successful login redirects to account dashboard</li>
                    </ol>
                  </div>

                  <div className="mb-4">
                    <h4>3. 💰 Managing Accounts</h4>
                    <ul>
                      <li><strong>View Accounts:</strong> See all your accounts and balances</li>
                      <li><strong>Create New Account:</strong> Support for 6 currencies</li>
                      <li><strong>Account Details:</strong> View transaction history</li>
                      <li><strong>Balance Updates:</strong> Real-time balance tracking</li>
                    </ul>
                  </div>

                  <div className="mb-4">
                    <h4>4. 🚀 Making Transfers</h4>
                    <ol>
                      <li>Go to "Transfer" section</li>
                      <li>Select your source account</li>
                      <li>Enter recipient account number</li>
                      <li>Specify amount and currency</li>
                      <li>Confirm transfer (instant processing)</li>
                    </ol>
                  </div>
                </Card.Body>
              </Card>
            </Col>
          </Row>
        </Tab>

        {/* Demo & Testing Tab */}
        <Tab eventKey="demo" title="Demo & Testing">
          <Row>
            <Col lg={10} className="mx-auto">
              <Alert variant="success" className="mb-4">
                <Alert.Heading>🎮 Try the Demo!</Alert.Heading>
                <p>
                  Perfect for recruiters and testers! Use the demo account to see transactions in action.
                </p>
              </Alert>

              <Card className="mb-4">
                <Card.Header><h3>🧪 Demo Account Testing</h3></Card.Header>
                <Card.Body>
                  <div className="mb-4">
                    <h4>Demo Account Information</h4>
                    <Alert variant="info">
                      <p><strong>Demo Account Number:</strong> <code>{demoAccountNumber}</code></p>
                      <p><strong>Purpose:</strong> Send money to this account to see automatic responses</p>
                      <p><strong>Auto-Response:</strong> You'll receive 2 demo transactions back</p>
                    </Alert>
                  </div>

                  <div className="mb-4">
                    <h4>How to Test:</h4>
                    <ol>
                      <li>Create your account and verify email</li>
                      <li>Use your $100 welcome credits</li>
                      <li>Send any amount to demo account: <code>{demoAccountNumber}</code></li>
                      <li>Watch for automatic demo transactions in return</li>
                      <li>Check your account balance and transaction history</li>
                    </ol>
                  </div>

                  <div className="mb-4">
                    <h4>Sample Test Scenarios:</h4>
                    <Table striped bordered>
                      <thead>
                        <tr>
                          <th>Test Case</th>
                          <th>Action</th>
                          <th>Expected Result</th>
                        </tr>
                      </thead>
                      <tbody>
                        <tr>
                          <td>Email Verification</td>
                          <td>Register new account</td>
                          <td>Receive verification email, $100 credit after verification</td>
                        </tr>
                        <tr>
                          <td>Currency Conversion</td>
                          <td>Create EUR account</td>
                          <td>Welcome amount shows as ~€95 (based on exchange rate)</td>
                        </tr>
                        <tr>
                          <td>Demo Transaction</td>
                          <td>Send $10 to demo account</td>
                          <td>Receive 2 automatic demo transactions back</td>
                        </tr>
                        <tr>
                          <td>Real-time Updates</td>
                          <td>Make any transfer</td>
                          <td>Balance updates immediately</td>
                        </tr>
                      </tbody>
                    </Table>
                  </div>
                </Card.Body>
              </Card>

              <Card>
                <Card.Header><h3>🔍 For Recruiters</h3></Card.Header>
                <Card.Body>
                  <p>This project demonstrates:</p>
                  <ul>
                    <li><strong>Full-Stack Development:</strong> Go backend + React frontend</li>
                    <li><strong>Database Design:</strong> PostgreSQL with proper relationships</li>
                    <li><strong>Security:</strong> JWT authentication, email verification</li>
                    <li><strong>Email Integration:</strong> SMTP with HTML templates</li>
                    <li><strong>Real-time Processing:</strong> Instant transaction updates</li>
                    <li><strong>Error Handling:</strong> Comprehensive error management</li>
                    <li><strong>UI/UX:</strong> Responsive design with modern aesthetics</li>
                    <li><strong>Testing:</strong> Automated demo responses for easy testing</li>
                  </ul>
                </Card.Body>
              </Card>
            </Col>
          </Row>
        </Tab>

        {/* API Reference Tab */}
        <Tab eventKey="api" title="API Reference">
          <Row>
            <Col>
              <Card className="mb-4">
                <Card.Header><h3>🔌 API Endpoints</h3></Card.Header>
                <Card.Body>
                  <h4>Authentication Endpoints</h4>
                  <Table striped>
                    <thead>
                      <tr><th>Method</th><th>Endpoint</th><th>Description</th></tr>
                    </thead>
                    <tbody>
                      <tr><td>POST</td><td>/v1/create_user</td><td>Create new user account</td></tr>
                      <tr><td>POST</td><td>/v1/login_user</td><td>User authentication</td></tr>
                      <tr><td>POST</td><td>/v1/verify_email</td><td>Email verification</td></tr>
                    </tbody>
                  </Table>

                  <h4>Account Management</h4>
                  <Table striped>
                    <thead>
                      <tr><th>Method</th><th>Endpoint</th><th>Description</th></tr>
                    </thead>
                    <tbody>
                      <tr><td>POST</td><td>/v1/accounts</td><td>Create new account</td></tr>
                      <tr><td>GET</td><td>/v1/accounts/:id</td><td>Get account details</td></tr>
                      <tr><td>GET</td><td>/v1/accounts</td><td>List user accounts</td></tr>
                    </tbody>
                  </Table>

                  <h4>Transactions</h4>
                  <Table striped>
                    <thead>
                      <tr><th>Method</th><th>Endpoint</th><th>Description</th></tr>
                    </thead>
                    <tbody>
                      <tr><td>POST</td><td>/v1/transfers</td><td>Create money transfer</td></tr>
                      <tr><td>GET</td><td>/v1/transfers</td><td>Get transfer history</td></tr>
                    </tbody>
                  </Table>
                </Card.Body>
              </Card>

              <Card>
                <Card.Header><h3>📊 Response Examples</h3></Card.Header>
                <Card.Body>
                  <h4>Successful User Creation</h4>
                  <pre className="bg-light p-3 rounded">
{`{
  "user": {
    "username": "john_doe",
    "full_name": "John Doe", 
    "email": "john@example.com",
    "created_at": "2025-01-01T12:00:00Z"
  }
}`}
                  </pre>

                  <h4>Account List Response</h4>
                  <pre className="bg-light p-3 rounded">
{`{
  "accounts": [
    {
      "id": 1,
      "owner": "john_doe",
      "balance": "10000",
      "currency": "USD",
      "created_at": "2025-01-01T12:00:00Z"
    }
  ]
}`}
                  </pre>
                </Card.Body>
              </Card>
            </Col>
          </Row>
        </Tab>
      </Tabs>

    </Container>
  );
};

export default Documentation;

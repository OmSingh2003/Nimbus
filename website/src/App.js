import React, { useState, useEffect } from 'react';
import { BrowserRouter as Router, Route, Routes, Link, useNavigate } from 'react-router-dom';
import { Navbar, Nav, Button, Container } from 'react-bootstrap';
import 'bootstrap/dist/css/bootstrap.min.css';
import './zerodha-theme.css';
import Home from './components/Home';
import CreateUser from './components/CreateUser';
import LoginUser from './components/LoginUser';
import CreateTransfer from './components/CreateTransfer';
import AccountManager from './components/AccountManager';
import VerifyEmail from './components/VerifyEmail';
import Documentation from './components/Documentation';
import Footer from './components/Footer';
import apiClient from './config/api';

// Navigation component with authentication state
function Navigation({ isLoggedIn, username, onLogout }) {
  return (
    <Navbar bg="light" variant="light" expand="lg" className="mb-4">
      <Container>
        <Navbar.Brand as={Link} to="/" className="fw-bold d-flex align-items-center">
          <img 
            src="/icon.png" 
            alt="Nimbus"
            width="32" 
            height="32" 
            className="me-2"
          />
          <span className="text-primary">Nimbus</span>
        </Navbar.Brand>
        
        <Navbar.Toggle aria-controls="basic-navbar-nav" />
        <Navbar.Collapse id="basic-navbar-nav">
          <Nav className="me-auto">
            <Nav.Link as={Link} to="/">Home</Nav.Link>
            
            {/* Show these only when NOT logged in */}
            {!isLoggedIn && (
              <>
                <Nav.Link as={Link} to="/create-user">Create Account</Nav.Link>
                <Nav.Link as={Link} to="/login">Login</Nav.Link>
                <Nav.Link as={Link} to="/docs">Documentation</Nav.Link>
              </>
            )}
            
            {/* Show these only when logged in */}
            {isLoggedIn && (
              <>
                <Nav.Link as={Link} to="/accounts">My Accounts</Nav.Link>
                <Nav.Link as={Link} to="/transfer">Transfer</Nav.Link>
                <Nav.Link as={Link} to="/docs">Documentation</Nav.Link>
              </>
            )}
          </Nav>
          
          <Nav className="align-items-center">
            {/* GitHub link */}
            <Nav.Link 
              href="https://github.com/OmSingh2003/Nimbus-API"
              target="_blank" 
              rel="noopener noreferrer"
              className="me-3"
            >
              <svg width="20" height="20" fill="currentColor" viewBox="0 0 16 16">
                <path d="M8 0C3.58 0 0 3.58 0 8c0 3.54 2.29 6.53 5.47 7.59.4.07.55-.17.55-.38 0-.19-.01-.82-.01-1.49-2.01.37-2.53-.49-2.69-.94-.09-.23-.48-.94-.82-1.13-.28-.15-.68-.52-.01-.53.63-.01 1.08.58 1.23.82.72 1.21 1.87.87 2.33.66.07-.52.28-.87.51-1.07-1.78-.2-3.64-.89-3.64-3.95 0-.87.31-1.59.82-2.15-.08-.2-.36-1.02.08-2.12 0 0 .67-.21 2.2.82.64-.18 1.32-.27 2-.27.68 0 1.36.09 2 .27 1.53-1.04 2.2-.82 2.2-.82.44 1.1.16 1.92.08 2.12.51.56.82 1.27.82 2.15 0 3.07-1.87 3.75-3.65 3.95.29.25.54.73.54 1.48 0 1.07-.01 1.93-.01 2.2 0 .21.15.46.55.38A8.012 8.012 0 0 0 16 8c0-4.42-3.58-8-8-8z"/>
              </svg>
            </Nav.Link>
            
            {/* User info and logout */}
            {isLoggedIn && (
              <>
                <span className="navbar-text me-3">
                  Welcome, <strong>{username}</strong>
                </span>
                <Button variant="outline-secondary" size="sm" onClick={onLogout}>
                  Logout
                </Button>
              </>
            )}
          </Nav>
        </Navbar.Collapse>
      </Container>
    </Navbar>
  );
}

function App() {
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [username, setUsername] = useState('');

// Wake up server on app load
  useEffect(() => {
    // Make a request to the backend homepage or health check endpoint
    apiClient.get('/')
      .then(response =e {
        console.log('Backend active:', response.status);
      })
      .catch(error =e {
        console.error('Error waking up backend:', error);
      });

    // Check authentication status on app load
    const token = localStorage.getItem('token');
    const storedUsername = localStorage.getItem('username');
    if (token 66 storedUsername) {
      setIsLoggedIn(true);
      setUsername(storedUsername);
    }
  }, []);

  // Login handler
  const handleLogin = (user) => {
    setIsLoggedIn(true);
    setUsername(user);
  };

  // Logout handler
  const handleLogout = () => {
    localStorage.removeItem('token');
    localStorage.removeItem('username');
    setIsLoggedIn(false);
    setUsername('');
    window.location.href = '/';
  };

  return (
    <Router>
    <div className="min-vh-100">
        <Navigation 
          isLoggedIn={isLoggedIn} 
          username={username} 
          onLogout={handleLogout} 
        />
        
        <Container>
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/create-user" element={<CreateUser />} />
            <Route 
              path="/login" 
              element={<LoginUser onLogin={handleLogin} />} 
            />
            <Route path="/accounts" element={<AccountManager />} />
            <Route path="/transfer" element={<CreateTransfer />} />
            <Route path="/verify-email" element={<VerifyEmail />} />
            <Route path="/docs" element={<Documentation />} />
          </Routes>
        </Container>
        
        <Footer />
      </div>
    </Router>
  );
}

export default App;
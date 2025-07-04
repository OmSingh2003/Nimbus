import React from 'react';
import { Container, Row, Col } from 'react-bootstrap';

const Footer = () => {
  return (
    <footer className="footer">
      <Container>
        <Row>
          <Col md={4}>
            <h5>Nimbus</h5>
            <p>
              Secure banking made simple. Experience modern financial services 
              with enterprise-grade security and user-friendly design.
            </p>
            <div className="d-flex">
              <a 
                href="https://github.com/OmSingh2003/Nimbus-API"
                target="_blank" 
                rel="noopener noreferrer"
                className="me-3"
              >
                GitHub
              </a>
              <a href="#" className="me-3">Documentation</a>
              <a href="#" className="me-3">API Reference</a>
            </div>
          </Col>
          
          <Col md={4}>
            <h5>Product</h5>
            <p><a href="/">Accounts</a></p>
            <p><a href="/">Transfers</a></p>
            <p><a href="/">Multi-Currency</a></p>
            <p><a href="/">Security</a></p>
          </Col>
          
          <Col md={4}>
            <h5>Support</h5>
            <p><a href="mailto:singhom2003.os@gmail.com">singhom2003.os@gmail.com</a></p>
            <p><a href="tel:+918810967896">+91 8810967896</a></p>
            <p><a href="/">Help Center</a></p>
          </Col>
        </Row>
        
        <div className="text-center">
          <p className="mb-0">
            © 2025 Nimbus. All rights reserved. Built with ❤️ for secure banking.
          </p>
        </div>
      </Container>
    </footer>
  );
};

export default Footer;

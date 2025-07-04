import React from 'react';
import { Card, Table } from 'react-bootstrap';
import apiClient from '../config/api';

const DebugInfo = () => {
  const token = localStorage.getItem('token');
  const username = localStorage.getItem('username');
  
  const debugInfo = {
    'API Base URL': apiClient.defaults.baseURL,
    'Environment': process.env.NODE_ENV,
    'API URL from ENV': process.env.REACT_APP_API_URL,
    'Token exists': !!token,
    'Token length': token ? token.length : 0,
    'Token preview': token ? `${token.substring(0, 20)}...` : 'None',
    'Username': username || 'None',
    'Current URL': window.location.href,
    'User Agent': navigator.userAgent.substring(0, 50) + '...'
  };

  return (
    <Card className="mt-3">
      <Card.Header>
        <h5>Debug Information</h5>
      </Card.Header>
      <Card.Body>
        <Table size="sm">
          <tbody>
            {Object.entries(debugInfo).map(([key, value]) => (
              <tr key={key}>
                <td><strong>{key}</strong></td>
                <td><code>{String(value)}</code></td>
              </tr>
            ))}
          </tbody>
        </Table>
        <hr />
        <h6>Test API Call</h6>
        <button 
          className="btn btn-sm btn-primary me-2"
          onClick={async () => {
            try {
              const response = await apiClient.get('/v1/accounts');
              console.log('Manual API test success:', response);
              alert('API call successful! Check console for details.');
            } catch (error) {
              console.error('Manual API test failed:', error);
              alert(`API call failed: ${error.response?.data?.message || error.message}`);
            }
          }}
        >
          Test /v1/accounts
        </button>
        <button 
          className="btn btn-sm btn-warning me-2"
          onClick={() => {
            localStorage.removeItem('token');
            localStorage.removeItem('username');
            alert('Local storage cleared! Please refresh the page and log in again.');
          }}
        >
          Clear Storage
        </button>
        <button 
          className="btn btn-sm btn-success"
          onClick={() => {
            window.location.reload();
          }}
        >
          Refresh Page
        </button>
      </Card.Body>
    </Card>
  );
};

export default DebugInfo;

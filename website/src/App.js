import React from 'react';
import { BrowserRouter as Router, Route, Routes, Link } from 'react-router-dom';
import 'bootstrap/dist/css/bootstrap.min.css';
import Home from './components/Home';
import CreateUser from './components/CreateUser';
import LoginUser from './components/LoginUser';
import CreateTransfer from './components/CreateTransfer';
import AccountManager from './components/AccountManager';

function App() {
  return (
    <Router>
      <div className="container">
        <nav className="navbar navbar-expand-lg navbar-light bg-light">
          <Link className="navbar-brand" to="/">VaultGuard</Link>
          <div className="collapse navbar-collapse">
            <ul className="navbar-nav mr-auto">
              <li className="nav-item">
                <Link className="nav-link" to="/">Home</Link>
              </li>
              <li className="nav-item">
                <Link className="nav-link" to="/create-user">Create User</Link>
              </li>
              <li className="nav-item">
                <Link className="nav-link" to="/login">Login</Link>
              </li>
              <li className="nav-item">
                <Link className="nav-link" to="/accounts">My Accounts</Link>
              </li>
              <li className="nav-item">
                <Link className="nav-link" to="/transfer">Transfer</Link>
              </li>
            </ul>
          </div>
        </nav>

        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/create-user" element={<CreateUser />} />
          <Route path="/login" element={<LoginUser />} />
          <Route path="/accounts" element={<AccountManager />} />
          <Route path="/transfer" element={<CreateTransfer />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;
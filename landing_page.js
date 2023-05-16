import React, { useState } from 'react';
import { AuthProvider, Descope } from '@descope/react';
import { Argyle } from 'argyle-sdk';
import './App.css';

function App() {
  const [isLoggedIn, setIsLoggedIn] = useState(false);

  const onSuccess = (e) => {
    console.log(e.detail.user.name);
    console.log(e.detail.user.email);
    setIsLoggedIn(true);
  };

  const onError = (err) => {
    console.log('Error!', err);
  };

  const argyleOptions = {
    token: 'your_argyle_token_here',
    environment: 'sandbox'
  };

  return (
    <div className="App">
      {isLoggedIn ? (
        <Argyle
          sdkOptions={argyleOptions}
          onLoad={() => console.log('Argyle component loaded')}
        />
      ) : (
        <header className="App-header">
          <nav className="navbar">
            <div className="navbar-container">
              <div className="navbar-brand">
                <h1 className="logo">Descope</h1>
              </div>
              <ul className="nav-list">
                <li className="nav-item">
                  <a href="#" className="nav-link">Features</a>
                </li>
                <li className="nav-item">
                  <a href="#" className="nav-link">Pricing</a>
                </li>
                <li className="nav-item">
                  <a href="#" className="nav-link">About</a>
                </li>
                <li className="nav-item">
                  <a href="#" className="nav-link">Contact</a>
                </li>
              </ul>
              <button className="btn btn-primary">Get Started</button>
            </div>
          </nav>
          <div className="hero-section">
            <h1 className="hero-title">Efficient project management made easy</h1>
            <p className="hero-description">Descope helps software development teams stay organized and on track.</p>
            <AuthProvider projectId="P2OLfBkdSWQCYRI18eTE6Y5qyVMn">
              <Descope
                flowId="sign-up-or-in"
                theme="light"
                onSuccess={onSuccess}
                onError={onError}
              />
            </AuthProvider>
          </div>
        </header>
      )}
      <footer className="App-footer">
        <p>&copy; 2023 Descope</p>
      </footer>
    </div>
  );
}

export default App;

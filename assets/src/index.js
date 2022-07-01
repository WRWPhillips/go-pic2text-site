import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';
import { BrowserRouter as Router } from 'react-router-dom';
import { applyMiddleware, createStore } from 'redux';
import { Provider } from 'react-redux';
import thunk from 'redux-thunk';

import Header from './components/Header';


// const store = createStore(rootReducers, applyMiddleware(thunk));
// window.store = store; //remove before complete deployment, here for development

const root = ReactDOM.createRoot(document.getElementById('root'));

root.render(
  // <Provider store={store}>
  
    <Router>
      <React.StrictMode>
        <Header />
        <App />
      </React.StrictMode>
    </Router>
  // </Provider>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();

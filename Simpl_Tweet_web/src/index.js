import React from 'react';
import ReactDOM from "react-dom";
import './index.css';
import App from './components/App';
import reportWebVitals from './reportWebVitals';
import {BrowserRouter} from "react-router-dom";

const rootElement = document.getElementById("root");
ReactDOM.render(
  //<React.StrictMode>
  //用这个包起来就可以用了
  <BrowserRouter>
      <App />
  </BrowserRouter>,
  //</React.StrictMode>
    rootElement
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();

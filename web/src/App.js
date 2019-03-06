import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';
import Login from './components/Login/Login'
import Home from './components/Home/Home'

class App extends Component {

  isLoggedIN(){
    return document.cookie.match("gs_session") != null
  }


  render() {
    return (
      <div >
        {this.isLoggedIN()?
        <Home/>:
        <Login/>
        
      }
      </div>
    );
  }
}

export default App;

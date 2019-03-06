import React, {Component} from 'react';
import './Login.css'
import tw from './tw.png'

class Login extends Component{

    render(){
        return(
            <div className='login-page'>
                <div className='login-form-cont'>
                    <h2>Gstreamer</h2>
                    <a href="http://http://142.93.145.74/auth">
                    <span className='login-btn'>
                        Login
                        <img src={tw} className='login-img'/>
                    </span>
                    </a>
                </div>
            </div>
        )
    }
}

export default Login;
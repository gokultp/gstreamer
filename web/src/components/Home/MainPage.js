import React,{Component} from 'react'
import axios from 'axios'
import Row from 'react-bootstrap/Row'
import Col from 'react-bootstrap/Col'

import Button from 'react-bootstrap/Button'
import './FirstPage.css'

class MainPage extends Component{
    state={
        events:[]
    }
    colors= ['aqua', 'black', 'blue', 'fuchsia', 'green', 
'lime', 'maroon', 'navy', 'olive', 'orange', 'purple', 'red', 
, 'teal', 'white', 'yellow'];

   
    
    componentDidMount(){
        var socket = new WebSocket("ws://142.93.145.74/socket");
        this.color=this.getColor(this.props.user.fav_streamer_name)
        console.log(this.color)
        socket.onmessage =  (e) =>{
            let data = JSON.parse(e.data)
            let events = [... data.data, ...this.state.events].slice(0, 20)
            this.setState({events})
        }
    }
    getColor(str){
        if(!str){
            return this.colors[5]
        }
        return this.colors[str.charCodeAt(0) * str.charCodeAt(1)% this.colors.length]
    }
    render(){
        let channel= this.props.user.fav_streamer_name
        return(
            <div>
                <Row>
                    <Col md={3} className='events'>
                        <h6>Events</h6>
                        <hr/>
                        {this.state.events.map(event=>(
                            <p><b style={{color: this.getColor(event.from_name)}}> @{event.from_name} </b> started following <b style={{color:this.color}}> @{event.to_name} </b></p>
                        ))}
                    </Col>
                    <Col md={6}>
                            <Row className='video'>
                                <center>
                                <iframe
                                    src={`https://player.twitch.tv/?channel=${channel}`}
                                    height="500"
                                    width="700"
                                    frameborder="0"
                                    scrolling="yes"
                                    allowfullscreen="yes">
                                </iframe>
                                </center>
                                
                            </Row>
                            
                    </Col>
                    <Col md={3}>
                        <iframe frameborder="0"
                            scrolling="yes"
                            src={`https://www.twitch.tv/embed/${channel}/chat`}
                            width="450"
                            height={window.innerHeight-20}>
                        </iframe>
                    </Col>
                </Row>
            </div>
            
            
        )
    }
}

export default MainPage;
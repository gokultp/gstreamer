import React,{Component} from 'react'
import axios from 'axios'
import FirstPage from './FIrstPage'
import Navbar from 'react-bootstrap/Navbar';
import MainPage from './MainPage'
import './FirstPage.css'
import Row from 'react-bootstrap/Row'
import Col from 'react-bootstrap/Col'
import Button from 'react-bootstrap/Button'
import Alert from 'react-bootstrap/Alert'

class Home extends Component{
    state={
        showAlert: false
    }
    componentDidMount(){
        this.getUser()
    }
    getUser(){
        axios.get('/user', {withCredentials: true}).then(res=>{
            console.log(res.data)
            this.setState({user: res.data})
        }).catch(err=>{
            console.log(err)
        })
    }
    resetStream(){
        var user = this.state.user;
        user.fav_streamer= null;
        this.setState({user})
    }

    onFavSubmit(favUser){
        axios.post("/user", {fav_streamer_name:favUser},{withCredentials: true}).then(res=>{
            this.setState({user: res.data})
        }).catch(err=>{
            this.setState({showAlert:true})

            setTimeout(()=>{
                this.setState({showAlert: false})
            }, 5000)
        })
    }
    render(){
        return(
            <div>
                 <Alert show={this.state.showAlert} variant="danger">
                    <Alert.Heading>Could not find that user</Alert.Heading>
                </Alert>

                <div >
                <Row className='nav'>
                    <Col md={8}>
                        <h3>gstreamer</h3>
                    </Col><Col md={2}>
                        <Button onClick={this.resetStream.bind(this)}variant="light">Change Streamer</Button>
                    </Col>
                    <Col md={2}>
                        <Button href="/logout">Logout</Button>
                    </Col>

                </Row>
                </div>
            {
                this.state.user?
                <div>
                    {this.state.user.fav_streamer?
                    <MainPage onChangeStream={this.resetStream.bind(this)} user={this.state.user}/>
                    :<FirstPage user={this.state.user} onSubmit={this.onFavSubmit.bind(this)}/>
                }
                </div>
                :null
            }
            </div>
            
            
        )
    }
}

export default Home;
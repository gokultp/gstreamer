import React,{Component} from 'react'
import './FirstPage.css'
import FormControl from 'react-bootstrap/FormControl'
import Row from 'react-bootstrap/Row'
import Col from 'react-bootstrap/Col'
import Button from 'react-bootstrap/Button';



class FirstPage extends Component{
    state={
        streamer: null

    }
    onChange(evt){
        this.setState({streamer: evt.target.value})
    }
    onSubmit(){
        this.props.onSubmit(this.state.streamer);
    }

    render(){
        const {logo, display_name} = this.props.user;
        const {streamer} = this.state;
        return(
            <div className='first-page'>
                <img src={logo}/>
                <h3>H!! {display_name}</h3>
                <div className='form-a'>
                    <Row>
                        <Col md={{ span: 3, offset: 3 }}>
                            <FormControl
                                value={streamer}
                                placeholder="Enter Fav streamer name to continue"
                                aria-describedby="basic-addon1"
                                onChange={this.onChange.bind(this)}
                            />
                        </Col>
                        <Col md={3}>
                            <Button onClick={this.onSubmit.bind(this)}>Next</Button>
                        </Col>
                    </Row>
                
                </div>
                
            </div>
        )
    }
}

export default FirstPage;
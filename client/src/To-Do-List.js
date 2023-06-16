import React,{Component} from "react";
import axois from "axiox"
import {Card, Header, Form, Input, Icon} from "sematic-ui-react";

let endpoint = "http://localhost:2020"


class ToDoList extends Component{

constructor(props){
    super(props);

    this.state={
        task:"",
        items:[],
    };
}

ComponentDidMount(){
    this.getask();
   }

render(){
return(
    <div>

        <div className="row">
            <header className="header" as="h2" color="blue" >
            TO DO LIST
            </header>    
        </div>

        <div className="row">
            <Form onSubmit={this.onSubmit }>

                <Input
                type ="text"
                name ="task"
                onChange ={this.onChange}
                value = {this.state.task}
                fluid
                placeholder ="Create Task"
                />
                 
            </Form>
        </div>

    </div>
);
}   

}


export default ToDoList
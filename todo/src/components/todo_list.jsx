import React, { Component } from 'react'
import Nav from './nav'
import { Modal, Button, Input, Row } from 'react-materialize'
import uuid from 'uuid'

export default class Todo extends Component {
    constructor(props) {
        super(props)
        this.state = {
            cards: null,
             title: '',
              body: '',
             date: ''
        }
    }
    async componentDidMount() {
        const fetchData = await fetch('api/get')
        const data = await fetchData.json()
        this.setState({cards: data})
    }
    render() {
        const { cards } = this.state
        if (!cards) return null
        
        return (
            <div>
            <Nav />
            <Modal
                header='Create a Todo'
                    trigger={<Button className="modalbtn">Create todo</Button>}>
                    <form onSubmit={this.handleSubmit}>
                        <Row>
    <Input placeholder="Title" s={12} label="Title" name="title" onChange={this.handleChange} />
</Row>
<Row>
  <Input s={12} placeholder="What to do?" type='textarea' name="body" onChange={this.handleChange} />
</Row>
<Row>
  <Input placeholder="Pick a time" name='on' type='date' name="date" onChange={this.handleChange} />
</Row>
<Button>Submit</Button>
</form>
             </Modal>
            <div className="contained">
            {this.renderCards()}
            </div>
             </div>
        )
    }
    renderCards = () => {
       const { cards } = this.state

       return cards.map(card => {
        return (
            <div className="acard" key={uuid()}>
            <h5>{card.title}</h5>
            <div className="content">
            <p>{card.body}</p>
            </div>
            </div>
        )
       })
    }
    handleChange = (e) => {
        this.setState({[e.target.name]: e.target.value})
    }
    handleSubmit = (e) => {
        e.preventDefault()
    }
}
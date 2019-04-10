import React, { Component } from 'react';
import './App.css';
import Nav from './Navigation'
import Table from './Table'
import CreateDialog from './CreateDialog'

class App extends Component {

  state = {
    headers: [],
    items: [],
    hideCreateDialog: true
  }

  selectOption(item) {
    if (item === 'New Deployment') {
      this.setState({
        hideCreateDialog: false
      })      
    } else {
      this.getItems(item)
    }
  }

  componentDidMount() {
    this.getItems('Pods')
  }

  getItems(item) {
    console.log(`Making call to fetch ${item}`)
    fetch(`/api/${item.toLowerCase().replace(/ /g, "")}`)
      .then(result => result.json())
      .then(result => {
        console.log('Result from server', result)
        this.setState({
          headers: result.Headers,
          items: result.Items
        })
      })
      .catch(e => console.error(e))
  }

  closeCreateDialog() {
    this.setState({
      hideCreateDialog: true
    })
  }

  render() {
    console.log('State', this.state)
    return (
      <div className="App">
        <div className="Nav">
          <Nav OnClick={this.selectOption.bind(this)} />
        </div>
        <div className="Table">
          <Table 
            headers={this.state.headers} 
            items={this.state.items} />
        </div>
        <div>
          <CreateDialog hide={this.state.hideCreateDialog} onClose={this.closeCreateDialog.bind(this)} />
        </div>
      </div>
    );
  }
}

export default App;

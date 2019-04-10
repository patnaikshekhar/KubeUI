import React from 'react'
import { Dialog, DialogType, DialogFooter } from 'office-ui-fabric-react/lib/Dialog'
import { PrimaryButton, DefaultButton } from 'office-ui-fabric-react/lib/Button'
import { Stack } from 'office-ui-fabric-react/lib/Stack'
import { TextField } from 'office-ui-fabric-react/lib/TextField'

class CreateDialog extends React.Component {

    state = {
        Name: '',
        Image: '',
        Namespace: 'default',
        Replicas: 1
    }

    createRecord() {
        fetch('/api/deployments', {
            headers: {
                'Content-Type': 'application/json',
            },
            method: 'POST',
            body: JSON.stringify(this.state)
        }).then(result => this.props.onClose())
          .catch(e => console.error(e))
    }     

    render() {

        return (
            <Dialog
                hidden={this.props.hide}
                dialogContentProps={{
                    type: DialogType.normal,
                    title: 'Create a new Deployment'
                }}>
            <Stack verticalAlign="true" tokens={{ childrenGap: 10 }}>
                <TextField value={this.state.Name} label="Name" onChange={(e) => this.setState({Name: e.target.value.toLowerCase().replace(/ /g, "-")})} />
                <TextField value={this.state.Image} label="Image" onChange={(e) => this.setState({Image: e.target.value})} />
                <TextField value={this.state.Namespace} label="Namespace" onChange={(e) => this.setState({Namespace: e.target.value})} />
                <TextField value={this.state.Replicas} label="Replicas" onChange={(e) => this.setState({Replicas: parseInt(e.target.value)})} />
            </Stack>
            <DialogFooter>
                <PrimaryButton onClick={this.createRecord.bind(this)} text="Save" />
                <DefaultButton onClick={this.props.onClose} text="Cancel" />
            </DialogFooter>
            </Dialog>
        )        
    }
}

export default CreateDialog
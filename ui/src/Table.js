import React from 'react'
import { DetailsList, DetailsListLayoutMode } from 'office-ui-fabric-react/lib/DetailsList'

const Table = (props) => {
    console.log(props.headers, props.items)
    return (
        <DetailsList
            items={props.items ? props.items : []}
            columns={props.headers ? props.headers.map(column => ({
                key: column,
                name: column,
                fieldName: column
            })): []}
            setKey="set"
            layoutMode={DetailsListLayoutMode.justified}
            isHeaderVisible={true}
            selectionPreservedOnEmptyClick={true}
            enterModalSelectionOnTouch={true}
          />
    )
}

export default Table
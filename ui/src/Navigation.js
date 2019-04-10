import React from 'react'
import { Nav } from 'office-ui-fabric-react/lib/Nav'

const Navigation = (props) => { 

    return (
        <Nav
            onLinkClick={(e, item) => props.OnClick(item.name)}
            expandedStateText="expanded"
            collapsedStateText="collapsed"
            selectedKey="key3"
            expandButtonAriaLabel="Expand or collapse"
            styles={{
                root: {
                width: 208,
                height: 350,
                boxSizing: 'border-box',
                border: '1px solid #eee',
                overflowY: 'auto'
                }
            }}
            groups={[
                {
                    links: [
                        {
                            name: 'New Deployment'
                        }
                    ]
                },
                {
                    links: [
                        {
                            name: 'Pods'
                        },
                        {
                            name: 'Deployments'
                        },
                        {
                            name: 'Services'
                        },
                        {
                            name: 'Jobs'
                        },
                        {
                            name: 'Cron Jobs'
                        }
                    ]
                }
            ]}
        />
    )
}

export default Navigation
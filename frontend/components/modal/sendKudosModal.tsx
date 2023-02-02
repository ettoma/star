import { Box, Button, Layer, Tag } from 'grommet'
import React from 'react'

function SendKudosModal({ setShow, clear, message, recipient }: { setShow: React.Dispatch<React.SetStateAction<boolean>>, clear: VoidFunction, message: string, recipient: string }) {
    return (
        <Layer
            background={{
                color: "rgba(20,20,20,0.95)"
            }}
            full={true}
            onEsc={() => {
                setShow(false)
                clear()
            }}
            onClickOutside={() => setShow(false)}
        >
            <Box justify="center" gap="large" pad="large" margin={{
                vertical: "200px",
                horizontal: "30px"

            }}>
                <Tag name="To" value={recipient} size='large' />
                <Tag name="Message" value={message} size='large' />
            </Box>
            <Button label="close" onClick={() => setShow(false)} margin="xlarge" />
        </Layer>
    )
}

export default SendKudosModal

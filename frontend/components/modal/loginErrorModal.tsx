import { Box, Button, Layer, Tag, Text } from 'grommet'
import React from 'react'

function ErrorModal({ setShowErrorModal, errorMessage }: { setShowErrorModal: React.Dispatch<React.SetStateAction<boolean>>, errorMessage: string }) {
    return (
        <Layer
            background={{
                color: "rgba(20,20,20,0.95)"
            }}
            full={true}
            onEsc={() => setShowErrorModal(false)}
            onClickOutside={() => setShowErrorModal(false)}
        >
            <Box justify="center" gap="large" pad="large" margin={{
                vertical: "200px",
                horizontal: "30px"

            }}>
                <Tag name='Error' value={errorMessage} size='medium' />
            </Box>
            <Button label="close" onClick={() => setShowErrorModal(false)} margin="xlarge" />
        </Layer>
    )
}

export default ErrorModal
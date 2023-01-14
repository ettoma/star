import { Box, Button, Heading, Paragraph } from "grommet"
import "react"
import { useNavigate } from "react-router"

function Home() {
    const navigate = useNavigate()
    return (
        <Box align="center" justify="center" color="active" gap="small">
            <Heading margin={{
                top: "xlarge",
                bottom: "large"
            }}>Welcome!</Heading>
            <Paragraph>Choose one of the options below:</Paragraph>
            <Button label="Login" primary onClick={() => navigate('/signin')} />
            <Button label="Register" onClick={() => navigate('/register')} />
        </Box>
    )
}

export default Home
import { Box, Button, Heading, PageHeader, Paragraph } from "grommet"
import "react"
import { useNavigate } from "react-router"

function Home() {
    const navigate = useNavigate()
    return (
        <Box justify="center" gap="large" pad="large">
            <PageHeader title="Welcome!" />
            <Paragraph>Choose one of the options below:</Paragraph>
            <Box gap="medium" align="center">
                <Button label="Login" primary onClick={() => navigate('/signin')} />
                <Button label="Register" onClick={() => navigate('/register')} />
            </Box>
        </Box>
    )
}

export default Home
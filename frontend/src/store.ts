import { configureStore, createListenerMiddleware } from '@reduxjs/toolkit'
import userReducer from '../reducers/userSlice'

const fetcher = createListenerMiddleware({})

fetcher.startListening({
    type: "SET",
    effect: () => console.log("set")
})


export default configureStore({
    reducer: {
        users: userReducer
    },
})
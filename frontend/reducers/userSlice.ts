import { createSlice } from '@reduxjs/toolkit'
import { handleGetUsers } from '../api/users/handleUsers'

export const userSlice = createSlice({
    name: 'users',
    initialState: {
        value: {},
    },

    //TODO : implement Reducer middleware to handle async call
    reducers: {
        getUsers: (state) => {
            handleGetUsers()
        },
    }
})

// Action creators are generated for each case reducer function
export const { getUsers } = userSlice.actions

export default userSlice.reducer
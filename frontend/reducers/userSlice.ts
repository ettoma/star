import { createSlice } from "@reduxjs/toolkit"
import { User } from "../api/models/user"

const initialState = {
    users: []


}

export const userSlice = createSlice({
    name: 'users',
    initialState: initialState,
    reducers: {
        setUserList: (state, action) => {
            state.users = action.payload
        }
    }
}
)

export const { setUserList } = userSlice.actions

export default userSlice.reducer
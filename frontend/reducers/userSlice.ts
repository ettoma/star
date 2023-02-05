import { createSlice } from "@reduxjs/toolkit"

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
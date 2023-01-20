import { createAsyncThunk, createSlice } from '@reduxjs/toolkit'
import { handleGetUsers } from '../api/users/handleUsers'



const initialState = {
    users: [
        {
            id: 1,
            name: "redux",
            username: "reduxTest",
            createdAt: 21321314128
        }
    ]
}
export const getUsers = createAsyncThunk(
    "getUsers",
    async (users, thunkAPI) => {
        try {
            const resp = await handleGetUsers().then((res) => res.json());

            return resp.data;
        } catch (error) {
            return thunkAPI.rejectWithValue('something went wrong');
        }
    }
);

export const userSlice = createSlice({
    name: 'users',
    initialState: initialState,

    reducers: {
        allUsers: (state) => {
            state.users
        },
    },
    extraReducers: (builder) => {
        builder
            .addCase(getUsers.pending, (state) => {
                console.log('loading');
            })
            .addCase(getUsers.fulfilled, (state, action) => {
                console.log(action.payload)
                state.users = action.payload;
            })
            .addCase(getUsers.rejected, (state, action) => {
                console.log(action);
            });
    },
})

export const { allUsers } = userSlice.actions

export default userSlice.reducer
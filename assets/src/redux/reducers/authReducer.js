import {
    LOGIN,
    SIGNUP,
    FETCHING,
    ERROR,
} from '../actions/authActions';

const initialState = {
    fetching: false,
    error: '',
    JWT: '',
}

const reducer = (state = initialState, action) => {
    switch (action.type) {
        case LOGIN:
            return {
                ...state,
                JWT: action.payload,
                fetching: false,
                error: '',
            }
        case FETCHING: 
            return {
                ...state,
                fetching: true,
            }
        case ERROR: 
            return {
                ...state,
                error: action.payload,
            }
        default: 
            return state;
    }
};

export default reducer;
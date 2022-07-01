import axios from 'axios';
import qs from 'qs';

export const LOGIN = "LOGIN";
export const SIGNUP = "SIGNUP";
export const FETCHING = "FETCHING";
export const ERROR = "ERROR";

const apiUrl = process.env.REACT_APP_API_URL;
function options(data) {
    return {
        method: 'POST',
        headers: {'content-type': 'application/x-www-url-form-encoded'},
        data: qs.stringify(data),
        }
}

export const login = (credentials) => async dispatch => {
    dispatch({
        type: FETCHING,
        payload: true,
    });
    try {
        console.log(credentials);
        const res = await axios.post(`${apiUrl}/api/signin`, options(credentials));

        localStorage.setItem({token: res.JWT});
    }   catch (error) {
        dispatch({
            type: ERROR,
            payload: error,
        });
    }
}
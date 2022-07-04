import axios from 'axios';
import qs from 'qs';

export const LOGIN = "LOGIN";
export const SIGNUP = "SIGNUP";
export const FETCHING = "FETCHING";
export const ERROR = "ERROR";

process.env.NODE_TLS_REJECT_UNAUTHORIZED = '0';

const apiUrl = process.env.REACT_APP_API_URL;
function options(data, route) {
    return {
        method: 'POST',
        headers: {'content-type': 'application/x-www-url-form-encoded'},
        data: qs.stringify(data),
        url: `${apiUrl}${route}`
        }
}

export const login = (credentials) => async dispatch => {
    dispatch({
        type: FETCHING,
        payload: true,
    });
    try {
        console.log(credentials);
        const res = await axios(options(credentials, '/api/signIn'));

        localStorage.setItem({token: res.JWT});
    }   catch (error) {
        dispatch({
            type: ERROR,
            payload: error,
        });
    }
}
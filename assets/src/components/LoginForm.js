import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { connect, useDispatch } from 'react-redux';

import { login } from '../redux/actions/authActions';

const LoginForm = props => {

    const navigate = useNavigate();
    const dispatch = useDispatch();
    
    const { error, fetching } = props;
    
    const [credentials, setCredentials] = useState({
        Username: '',
        Password: '',
    });

    const handleChanges = e => {
        setCredentials({
            ...credentials,
            [e.target.name]: e.target.value,
        });
    };

    const handleSubmit = e => {
        e.preventDefault();
        dispatch(login(credentials));
        navigate('/upload');
    }

    return (
    <div>
        <h3>Enter login details here </h3>
        <form 
        onChange={handleChanges}
        >
            <label>
                Name:
            <input type="string" name="Username" />
            </label>
            <label>
                Password:
            <input type="password" name="Password" />
            </label>
            <input type="submit" value="submit" onClick={handleSubmit}/>
        </form>
    </div>
    );
};

const mapStateToProps = state => {
    return {
        auth: state.authReducer,
    };
};


export default connect(mapStateToProps, { login })(LoginForm);
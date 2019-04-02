import React from 'react';
import { BuildForm, formToJSON } from '../utils';

const submitLoginFunc = (submitter) => (ev) => {
	ev.preventDefault();
	const form = ev.target;
	submitter(formToJSON(form));
	return false;
}//-- end submitLoginFunc

const LoginForm = BuildForm({
	'Username': { 'type': 'text', 'required': true },
	'Password': { 'type': 'password', 'required': true },
});

/*
const LoginForm = ({ submitLogin }) => (
	<form id="login-form" onSubmit={submitLoginFunc(submitLogin)}>
		<label htmlFor="Username">Username:</label>
		<input name="Username" type="text" />
		<label htmlFor="Password">Password:</label>
		<input name="Password" type="password" />
		<button type="submit">Login</button>
	</form>
)//-- end LoginForm
*/

export default LoginForm;


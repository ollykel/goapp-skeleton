import fetch from 'cross-fetch';
import { postFetch } from './utils.js';

export const SUBMIT_LOGIN = 'SUBMIT_LOGIN';
export const CONFIRM_LOGIN = 'CONFIRM_LOGIN';
export const CHECK_LOGIN = 'CHECK_LOGIN';
export const SET_ACCOUNT_INFO = 'SET_ACCOUNT_INFO';
export const SUBMIT_ACCOUNT = 'SUBMIT_ACCOUNT';
export const CONFIRM_ACCOUNT = 'CONFIRM_ACCOUNT';

export const SubmitLogin = () => ({
	type: 'SUBMIT_LOGIN',
})//-- end SubmitLogin

export const ConfirmLogin = ({ Authorized }) => ({
	type: 'CONFIRM_LOGIN',
	Authorized
})//-- end ConfirmLogin

const SubmitAccount = () => ({
	type: 'SUBMIT_ACCOUNT',
})//-- end SubmitAccount

const ConfirmAccount = ({ Success, Error }) => ({
	type: 'CONFIRM_ACCOUNT',
	Success,
	Error
})//-- end ConfirmAccount

export const CreateAccount = ({ Username, Email,
		Password }) => (dispatch) => {
	dispatch(SubmitAccount());
	return postFetch('/api/account',
		(confirmation) => dispatch(ConfirmAccount(confirmation)));	
}//-- end CreateAccount

export const Login = ({ Username, Password }) => (dispatch) => {
	dispatch(SubmitLogin());
	return fetch('/api/login/', {
		method: 'POST',
		body: JSON.stringify({ Username, Password }),
		headers: {
			'Accept': 'application/json',
			'Content-Type': 'application/json'
		}
	}).then((resp) => resp.json()).then(({ Success }) => {
		dispatch(ConfirmLogin({ Authorized: Success }))
		if (Success) dispatch(GetAccountInfo());
	});
}//-- end Login



import React from 'react';
import fetch from 'cross-fetch';

export const formToJSON = (form) => {
	let data = new FormData(form), elem = {}, output = {};
	for (let key of data.keys()) {
		elem = form.elements[key];
		output[key] = elem.type === 'file' ? elem.files : elem.value;
	}//-- end for pair
	return output 
}//-- end const formToJSON

export const OnSubmit = (submitter) => (ev) => {
	ev.preventDefault();
	submitter(formToJSON(ev.target));
	return false;
}//-- end const OnSubmit

const parseFieldDefault = ({ key, name, type, required, value }) => [
	<label key={key} htmlFor={name}>{name}</label>,
	<input key={key + 1} name={name} type={type} required={required}
		value={value} placeholder={name} />
]//-- end const parseField 

const parseFile = ({ key, name, required, accept }) => [
	<label key={key} htmlFor={name}>{name}</label>,
	<input key={key + 1} name={name} type='file' required={required}
		accept={accept} />
]//-- end parseFile

const parseTextarea = ({ key, name, required, value, rows, cols }) => [
	<label key={key} htmlFor={name}>{name}</label>,
	<textarea key={key + 1} name={name} required={required} rows={rows}
		cols={cols} placeholder={name}>{value}</textarea>
]//-- end const parseTextarea

const parseField = (qualities) => {
	switch (qualities.type) {
		case 'textarea':
			return parseTextarea(qualities);
		case 'file':
			return parseFile(qualities);
		default:
			return parseFieldDefault(qualities);
	}//-- end switch
}//-- end parseField

export const BuildForm = (fields) => ({ submitter }) => {
	return (
		<form onSubmit={OnSubmit(submitter)}>
			{ Object.keys(fields).map((name, k) => parseField({
				key: k * 2, name, ...fields[name]
			})) }
			<button type="Submit">Submit</button>
		</form>
	);//-- end return
}//-- end const BuildForm

const methodFetch = (method) => (path, onSuccess) => (data) => {
	console.log('%s data: %o', method, data);
	return fetch(path, {
		method,
		body: JSON.stringify(data),
		headers: {
			'Accept': 'application/json',
			'Content-Type': 'application/json'
		}
	}).then((resp) => resp.json()).then((body) => onSuccess(body));
}//-- end postFetch

export const postFetch = methodFetch('POST');
export const putFetch = methodFetch('PUT');
export const deleteFetch = methodFetch('DELETE');

const matchConfirmPassword = (accountData) => {
	if (accountData['Password'] !== accountData['ConfirmPassword']) {
		return '"Password" must match "Confirm Password"';
	}
	return '';
}//-- end matchConfirmPassword

const checkPasswordLength = (accountData) => {
	let password = accountData['Password'];
	if (password.length < minPasswordLength) {
		return `Password must be at least ${minPasswordLength} chars long`;
	}
	return '';
}//-- end checkPasswordLength

// list of functions to determine whether password is valid given account
// data
const passwordValidators = [ matchConfirmPassword, checkPasswordLength ];

const validatePassword = (accountData) => {
	let err = '';
	for (let validator of passwordValidators) {
		err = validator(accountData);
		if (err !== '') return err;
	}//-- end for validator
	return ''
}//-- end const validatePassword


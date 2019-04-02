import { combineReducers } from 'redux';

const Root = (state = {
	Authorized: false
}, action) => {
	if (action.hasOwnProperty('Authorized')) {
		const { Authorized } = action;
		return { ...state, Authorized };
	}
	return state;
}//-- end Root

const LoginPage = (state = {
	hasSubmitted: false,
	isSubmitting: false,
	Authorized: false
}, action) => {
	switch (action.type) {
		case SUBMIT_LOGIN:
			return { ...state, hasSubmitted: true, isSubmitting: true };
		case CONFIRM_LOGIN:
			const { Authorized } = action;
			return { ...state, isSubmitting: false, Authorized };
		default:
			return { ...state, hasSubmitted: false, Authorized: false };
	}//-- end switch
}//-- end LoginPage

const AccountPage = (state = {
	hasSubmitted: false,
	isSubmitting: false,
	Success: false,
	Error: ''
}, action) => {
	switch (action.type) {
		case SUBMIT_ACCOUNT:
			return { ...state, hasSubmitted: true, isSubmitting: true };
		case CONFIRM_ACCOUNT:
			const { Success, Error } = action;
			return { ...state, isSubmitting: false, Success, Error };
		default:
			return state;
	}//-- end switch
}//-- end AccountPage

const rootReducer = combineReducers({
	Root,
	LoginPage,
	AccountPage
});

export default rootReducer;


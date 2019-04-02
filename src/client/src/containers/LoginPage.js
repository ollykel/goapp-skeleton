import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { Link } from 'react-router-dom';
import { Redirect } from 'react-router';
import { connect } from 'react-redux';
import LoginForm from '../components/LoginForm';
import { Notice } from '../components/standard';
import { Login } from '../actions';

class LoginPage extends Component {
	constructor(props) {
		super(props);
		this.submitLogin = this.submitLogin.bind(this);
	}//-- end constructor

	submitLogin({ Username, Password }) {
		const { dispatch } = this.props;
		dispatch(Login({ Username, Password }));
	}//-- end submitLogin

	render() {
		const { isSubmitting, hasSubmitted, Authorized } = this.props;
		return (
			<div id="login-page">
				{Authorized ? <Redirect to="/" /> : ''}
				<Link to="/account">Create Account</Link>
				<h2>{"Please log in:"}</h2>
				{isSubmitting ? Notice('Please wait...') : ''}
				{hasSubmitted && !Authorized ? Notice('Login failed') : ''}
				<LoginForm submitter={this.submitLogin} />
			</div>
		);//-- end return
	}//-- end render
}//-- end class LoginPage

LoginPage.propTypes = {
	isSubmitting: PropTypes.bool.isRequired,
	hasSubmitted: PropTypes.bool.isRequired,
	Authorized: PropTypes.bool.isRequired
};//-- end propTypes

const mapStateToProps = ({ LoginPage }) => LoginPage

export default connect(mapStateToProps)(LoginPage);


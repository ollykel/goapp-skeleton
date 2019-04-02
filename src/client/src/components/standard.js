import React from 'react';
import { Link } from 'react-router-dom';

export const Notice = (msg) => (<span><strong>{msg}</strong></span>)

const breadcrumb = ({ key, name, path }) => [
	<span key={key}> / </span>,
	<Link key={key + 1} to={path}>{name}</Link>
]

const buildCrumbs = (path) => {
	let crumbs = path.map((name) => ({ name }));
	let url = '';
	for (let i in crumbs) {
		url += '/' + crumbs[i].name;
		crumbs[i].path = url;
	}//-- end for i
	crumbs[0] = { name: 'home', path: '/' };
	return crumbs.map((data, i) => breadcrumb({ ...data, key: i * 2 }));
}//-- end buildCrumbs

export const Breadcrumbs = ({ path }) => (
	<div className="breadcrumbs">{buildCrumbs(path)}</div>
)//-- end Breadcrumbs

const submitDelete = (onDelete) => (ev) => {
	ev.preventDefault();
	onDelete(ev);
	return false;
}//-- end submitDelete

export const DeleteButton = ({ onDelete }) => (
	<a href="#delete" onClick={submitDelete(onDelete)}>Delete</a>
)//-- end DeleteButton


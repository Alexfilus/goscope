package goscope_css

func WatcherStyles() string {
	const styles = `
	h1, h2, h3 {
		font-family: "Manrope", "sans-serif";
	}
	body {
		background: #1c1c1c;
		color: #fff;
		font-family: "Manrope", "sans-serif";
	}

	.monospaced {
		font-family: "Fira Mono", monospace
	}

	.navbar-link {
		text-decoration: none;
		color: #fff;
		border: none;
	}
	.active-navbar-link {
		text-decoration: none;
		color: #7cc3bf;
	}
	.navbar-link:hover {
		text-decoration: underline;
		color: #7cc3bf;
	}
	.custom-td {
		background: #2B2B2D;
	}

	th {
		border-radius: 4px;
	}
	.eye-image {
		fill: #fff;
	}
	.eye-image:hover {
		fill: #7cc3bf;
	}
	table {
		border-spacing: 2px;
		border-collapse: separate;
	}

	.paginate-button {
		color: #fff;
		border: none;
	}
	.paginate-button:hover {
		color: #7cc3bf;
		text-decoration: underline;
	}
	.tab-button {
		color: #fff;
		border: none;
		background: #2B2B2D;
		border-radius: 4px;
	}
	.badge-success {
		padding: 4px;
		border-radius: 4px;
		color: #fff;
		background-color: #1f9d55;
	}
	.badge-info {
		padding: 4px;
		border-radius: 4px;
		color: #fff;
		background-color: #1c3d5a;
	}
	.badge-turq {
		padding: 4px;
		border-radius: 4px;
		color: #fff;
		background-color: #503194;
	}
	.badge-secondary {
		border-radius: 4px;
		padding: 4px;
		color: #fff;
		background-color: #494444;
	}
	.badge-danger {
		border-radius: 4px;
		padding: 4px;
		color: #fff;
		background-color: #621b18;
	}
	.badge-warning {
		border-radius: 4px;
		padding: 4px;
		color: #fff;
		background-color: #684f1d;
	}
	`
	return MinifyCss(styles)
}

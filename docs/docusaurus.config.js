const config = {
  title: 'Nine Self Service API',
  favicon: 'favicon.ico',
  url: 'https://docs.nineapis.ch',
  baseUrl: '/',
  presets: [
    [
      '@docusaurus/preset-classic',
      {
        docs: false,
      },
    ],
    [
      'redocusaurus',
      {
        specs: [
          {
            spec: 'openapi-spec.json',
            route: '/',
          },
        ],
      },
    ]
  ],
  themeConfig: {
    navbar: {
      title: 'API docs',
      logo: {
        alt: 'Nine Logo',
        src: 'logo.svg',
      },
      items: [
        {
          title: 'GitHub',
          href: 'https://github.com/ninech/apis',
          label: 'GitHub',
          position: 'right',
        }
      ],
    },
    colorMode: {
      defaultMode: 'light',
      disableSwitch: false,
      // TODO: enable respect color scheme once this bug is fixed:
      // https://github.com/rohit-gohri/redocusaurus/issues/215
      respectPrefersColorScheme: false,
    },
  },
}

module.exports = config;

from setuptools import setup, find_packages

__version__ = "0.1"

install_requires = [
    'mkdocs', 'mkdocs-material'
]

setup(
    name="etherniti-docs",
    version=__version__,
    packages=find_packages(),
    install_requires=install_requires,
    package_data={
        '': ['*.rst'],
    },
    author="etherniti",
    author_email="team@etherniti.og",
    description="Documentation files for etherniti",
    license="GPL3",
    keywords="docker etherniti docs",
    url="https://github.com/etherniti",
    classifiers=[
        'Development Status :: 4 - Beta',
        'Environment :: Other Environment',
        'Intended Audience :: Developers',
        'Operating System :: OS Independent',
        'Programming Language :: Python',
        'Programming Language :: Python :: 2.7',
        'Programming Language :: Python :: 3.4',
        'Programming Language :: Python :: 3.5',
        'Programming Language :: Python :: 3.6',
        'Topic :: Utilities',
    ],
)

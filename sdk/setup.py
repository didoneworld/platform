"""SDK Package Setup"""
from setuptools import setup, find_packages

setup(
    name="didoneworld-sdk",
    version="0.1.0",
    description="DID One World - Universal Identity SDK",
    packages=find_packages(),
    install_requires=["requests>=2.28.0"],
    python_requires=">=3.9",
    extras_require={
        "dev": ["pytest>=7.0.0"],
    },
)

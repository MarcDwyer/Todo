import React from 'react'
import { NavItem, Navbar } from 'react-materialize'


const Nav = () => (
    <Navbar brand='logo' right className="amber darken-1">
        <NavItem onClick={() => console.log('test click')}>Getting started</NavItem>
        <NavItem href='#'>Components</NavItem>
</Navbar>
)

export default Nav;
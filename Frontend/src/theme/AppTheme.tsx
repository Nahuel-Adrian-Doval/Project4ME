import { ThemeProvider } from "@emotion/react"
import { CssBaseline } from "@mui/material"
import {theme} from './'
import React from "react"


const AppTheme = ({ children }: { children: React.ReactNode }) => {
  return (
    <ThemeProvider theme={theme}>
        <CssBaseline />
        {children}
    </ThemeProvider>
  )
}

export default AppTheme
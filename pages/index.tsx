import Image from 'next/image'

import Box from '@material-ui/core/Box';
import Card from '@material-ui/core/Card';
import CardActionArea from '@material-ui/core/CardActionArea';
import CardHeader from '@material-ui/core/CardHeader';
import Container from '@material-ui/core/Container';
import CssBaseline from '@material-ui/core/CssBaseline';
import Grid from '@material-ui/core/Grid';
import Hidden from '@material-ui/core/Hidden';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableRow from '@material-ui/core/TableRow';
import Typography from '@material-ui/core/Typography';

import anthosLogo from '../public/anthos.png';
import diffyLogo from '../public/diffy.png';
import mapsLogo from '../public/maps.png';
import profilePic from '../public/profile.jpg';
import styles from './index.module.css';

function ProfileHeader() {
    return (
        <Grid container alignItems="center">
            <Grid item xs={8} md={10}>
                <Typography variant="h4" component="h1">
                    Masaya Suzuki (draftcode)
                </Typography>
                <Typography variant="body1" className={styles.subtitle}>
                    Software Engineer. SF Bay Area. Cat lover.
                </Typography>
            </Grid>
            <Grid item xs={4} md={2}><Image src={profilePic} alt="face picture" /></Grid>
        </Grid>
    )
}

function WorkExperience() {
    return (
        <Grid container spacing={2}>
            <Grid item>
                <Typography variant="h5" component="h2">
                    Work Experience
                </Typography>
            </Grid>
            <Grid item>
                <Card variant="outlined">
                    <CardHeader
                        title="Google, LLC"
                        subheader="Software Engineer, August 2015&ndash;"
                    />
                    <Box className={styles.cardContent}>
                        <Grid container spacing={1}>
                            <Hidden xsDown>
                                <Grid item xs={1}><Image src={anthosLogo} alt="Anthos logo" /></Grid>
                            </Hidden>
                            <Grid item xs={12} sm={11}>
                                <Typography paragraph variant="body1">
                                    <b>Anthos</b>: Provide managed Kubernetes clusters in multi-cloud.
                                </Typography>
                                <Typography paragraph variant="body1">
                                    Tech lead. The role is to coordinate with internal teams to provide managed
                                    Kubernetes clusters on AWS and Azure.
                                </Typography>
                            </Grid>
                            <Hidden xsDown>
                                <Grid item xs={1}><Image src={diffyLogo} alt="Diffy logo" /></Grid>
                            </Hidden>
                            <Grid item xs={12} sm={11}>
                                <Typography paragraph variant="body1">
                                    <b>googlesource.com</b>: Git server for Google products
                                </Typography>

                                <Typography paragraph variant="body1">
                                    Tech lead. The role was to lead the team to provide Git repositories used by
                                    Chromium, Android, etc.. See also the activities related to this.
                                </Typography>
                            </Grid>
                            <Hidden xsDown>
                                <Grid item xs={1}><Image src={mapsLogo} alt="Google Maps logo" /></Grid>
                            </Hidden>
                            <Grid item xs={12} sm={11}>
                                <Typography paragraph variant="body1">
                                    <b>Google Maps</b>: API frontend for Google Maps
                                </Typography>

                                <Typography paragraph variant="body1">
                                    The role was to create the server infrastructure for Google Maps.
                                </Typography>
                            </Grid>
                        </Grid>
                    </Box>
                </Card>
            </Grid>
            <Grid item>
                <Card variant="outlined">
                    <CardHeader
                        title="Google Japan Inc."
                        subheader="Software Engineer, April 2014&ndash;August 2015"
                    />
                    <Box className={styles.cardContent}>
                        <Grid container spacing={1}>
                            <Hidden xsDown>
                                <Grid item xs={1}><Image src={mapsLogo} alt="Google Maps logo" /></Grid>
                            </Hidden>
                            <Grid item xs={12} sm={11}>
                                <Typography paragraph variant="body1">
                                    <b>Google Maps</b>: API frontend for Google Maps
                                </Typography>

                                <Typography paragraph variant="body1">
                                    The role was to create the server infrastructure for Google Maps. Transferred to the
                                    headquarter after an year.
                                </Typography>
                            </Grid>
                        </Grid>
                    </Box>
                </Card>
            </Grid>
        </Grid>
    )
}

function Activities() {
    return (
        <Grid container spacing={2}>
            <Grid item xs={12}>
                <Typography variant="h5" component="h2">
                    Activities
                </Typography>
            </Grid>
            <Grid item xs={12} sm={6}>
                <Card variant="outlined">
                    <CardActionArea href="https://github.com/google/hprof-parser">
                        <CardHeader
                            title="hprof-parser"
                            subheader="JVM heap dump parser"
                        />
                    </CardActionArea>
                </Card>
            </Grid>
            <Grid item xs={12} sm={6}>
                <Card variant="outlined">
                    <CardActionArea href="https://github.com/google/goblet">
                        <CardHeader
                            title="goblet"
                            subheader="Git caching proxy"
                        />
                    </CardActionArea>
                </Card>
            </Grid>
            <Grid item xs={12} sm={6}>
                <Card variant="outlined">
                    <CardActionArea href="https://github.com/google/gitprotocolio">
                        <CardHeader
                            title="gitprotocolio"
                            subheader="A Git protocol parser written in Go"
                        />
                    </CardActionArea>
                </Card>
            </Grid>
            <Grid item xs={12} sm={6}>
                <Card variant="outlined">
                    <CardActionArea href="https://github.com/google/ijaas">
                        <CardHeader
                            title="ijaas"
                            subheader="IntelliJ as a Service"
                        />
                    </CardActionArea>
                </Card>
            </Grid>
        </Grid>
    )
}

function Education() {
    return (
        <Grid container spacing={2}>
            <Grid item>
                <Typography variant="h5" component="h2">
                    Education
                </Typography>
            </Grid>
            <Grid item>
                <Card variant="outlined">
                    <CardHeader
                        title="MS Computer Science Tokyo Institute of Technology"
                        subheader="March 2014. Takuo Watanabe Lab."
                    />
                    <Box className={styles.cardContent}>
                        <Table>
                            <TableBody>
                                <TableRow>
                                    <TableCell align="right">
                                        Concentration
                                    </TableCell>
                                    <TableCell>
                                        Model checking and Fault tolerance
                                    </TableCell>
                                </TableRow>
                                <TableRow>
                                    <TableCell align="right" style={{verticalAlign: "top", borderBottom: "0"}}>
                                        Thesis
                                    </TableCell>
                                    <TableCell style={{borderBottom: "0"}}>
                                        <Typography paragraph variant="body2">
                                            Full-Automatic Exhaustive Fault-Injection on Software Models of Message-Passing Systems
                                        </Typography>

                                        <Typography paragraph variant="body2">
                                            Fault tolerance of distributed systems can be effectively verified by model
                                            checking and fault injection, but its process is highly error-prone. I proposed
                                            a way to solve this problem by adding a language support to modeling languages,
                                            which is a common approach in programming languages.
                                        </Typography>
                                    </TableCell>
                                </TableRow>
                            </TableBody>
                        </Table>
                    </Box>
                </Card>
            </Grid>
            <Grid item>
                <Card variant="outlined">
                    <CardHeader
                        title="BS Computer Science Tokyo Institute of Technology"
                        subheader="March 2012. Takuo Watanabe Lab."
                    />
                    <Box className={styles.cardContent}>
                        <Table>
                            <TableBody>
                                <TableRow>
                                    <TableCell align="right">
                                        Concentration
                                    </TableCell>
                                    <TableCell>
                                        Context-oriented programming
                                    </TableCell>
                                </TableRow>
                                <TableRow>
                                    <TableCell align="right" style={{verticalAlign: "top", borderBottom: "0"}}>
                                        Thesis
                                    </TableCell>
                                    <TableCell style={{borderBottom: "0"}}>
                                        <Typography paragraph variant="body2">
                                            An Implementation Method of Context-Oriented Programming in Objective-C<br />
                                        </Typography>

                                        <Typography paragraph variant="body2">
                                            Context-oriented programming is a programming method that enables us to define behaviors
                                            that depend on the program&apos;s execution context. I proposed an implementation method of
                                            Context-oriented programming in Objective-C.
                                        </Typography>
                                    </TableCell>
                                </TableRow>
                            </TableBody>
                        </Table>
                    </Box>
                </Card>
            </Grid>
        </Grid>
    )
}

export default function Home() {
    return (
        <>
            <CssBaseline />
            <Container maxWidth="md" className={styles.root}>
                <ProfileHeader />
                <WorkExperience />
                <Activities />
                <Education />
                <div>
                    <ul>
                        <li>
                            <a href='https://github.com/draftcode'>https://github.com/draftcode</a>
                        </li>
                        <li>
                            <a href='https://twitter.com/draftcode'>https://twitter.com/draftcode</a>
                        </li>
                        <li>
                            <a href='https://draftcode.osak.jp'>https://draftcode.osak.jp</a>
                        </li>
                    </ul>
                </div>
            </Container>
        </>
    )
}

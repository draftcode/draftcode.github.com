import type { ReactNode } from 'react';
import anthosLogo from './anthos.png';
import aviatorLogo from './aviator.jpg';
import diffyLogo from './diffy.png';
import launchableLogo from './launchable.jpg';
import mapsLogo from './maps.png';
import profilePic from './profile.jpg';

export default function App() {
  return (
    <div className="container mx-auto my-4 space-y-8">
      <ProfileHeader />
      <WorkExperience />
      <Activities />
      <Education />
      <div>
        <ul>
          <li>
            <a href="https://github.com/draftcode">
              https://github.com/draftcode
            </a>
          </li>
          <li>
            <a href="https://twitter.com/draftcode">
              https://twitter.com/draftcode
            </a>
          </li>
          <li>
            <a href="https://draftcode.osak.jp">https://draftcode.osak.jp</a>
          </li>
        </ul>
      </div>
    </div>
  );
}

function ProfileHeader() {
  return (
    <div className="flex justify-between items-center">
      <div>
        <h1 className="text-3xl">Masaya Suzuki (draftcode)</h1>
        <p>Software Engineer. SF Bay Area. Cat lover.</p>
      </div>
      <img className="size-64" src={profilePic} alt="face" />
    </div>
  );
}

function Card({
  title,
  subtitle,
  children,
}: {
  title: string;
  subtitle: string;
  children?: ReactNode;
}) {
  return (
    <div className="border-2 p-2 rounded">
      <h3 className="text-xl">{title}</h3>
      <p className="text-gray-500">{subtitle}</p>
      {children && <div className="p-2 space-y-8">{children}</div>}
    </div>
  );
}

function WorkExperience() {
  return (
    <div>
      <h2 className="text-2xl my-4">Work Experience</h2>
      <div className="space-y-2">
        <Card
          title="Aviator Technologies, Inc."
          subtitle="Head of Engineering, February 2023&ndash;"
        >
          <div className="flex gap-x-2">
            <img className="size-20" src={aviatorLogo} alt="Aviator logo" />
            <div className="space-y-4">
              <p>
                <b>Aviator</b>: A suite of developer productivity tools,
                inspired by Google.
              </p>
            </div>
          </div>
        </Card>
        <Card
          title="Launchable, Inc."
          subtitle="Principal Software Engineer, August 2021&ndash;February 2023"
        >
          <div className="flex gap-x-2">
            <img
              className="size-20"
              src={launchableLogo}
              alt="Launchable logo"
            />
            <div className="space-y-4">
              <p>
                <b>Predictive Test Selection</b>: Predict which tests are more
                likely to fail by using Machine Learning.
              </p>
            </div>
          </div>
        </Card>
        <Card
          title="Google, LLC"
          subtitle="Senior Software Engineer, August 2015&ndash;August 2021"
        >
          <div className="flex gap-x-2">
            <img className="size-20" src={anthosLogo} alt="Anthos logo" />
            <div className="space-y-4">
              <p>
                <b>Anthos</b>: Provide managed Kubernetes clusters in
                multi-cloud.
              </p>
              <p>
                Tech lead. The subtitle is to coordinate with internal teams to
                provide managed Kubernetes clusters on AWS and Azure.
              </p>
            </div>
          </div>
          <div className="flex gap-x-2">
            <img className="size-20" src={diffyLogo} alt="Diffy logo" />
            <div className="space-y-4">
              <p>
                <b>googlesource.com</b>: Git server for Google products
              </p>
              <p>
                Tech lead. The subtitle was to lead the team to provide Git
                repositories used by Chromium, Android, etc.. See also the
                activities related to this.
              </p>
            </div>
          </div>
          <div className="flex gap-x-2">
            <img className="size-20" src={mapsLogo} alt="Google Maps logo" />
            <div className="space-y-4">
              <p>
                <b>Google Maps</b>: API frontend for Google Maps
              </p>
              <p>
                The subtitle was to create the server infrastructure for Google
                Maps.
              </p>
            </div>
          </div>
        </Card>
        <Card
          title="Launchable, Inc."
          subtitle="Software Engineer, April 2014&ndash;August 2015"
        >
          <div className="flex gap-x-2">
            <img className="size-20" src={mapsLogo} alt="Google Maps logo" />
            <div className="space-y-4">
              <p>
                <b>Google Maps</b>: API frontend for Google Maps
              </p>

              <p>
                The subtitle was to create the server infrastructure for Google
                Maps. Transferred to the headquarter after an year.
              </p>
            </div>
          </div>
        </Card>
      </div>
    </div>
  );
}

function Activities() {
  return (
    <div>
      <h2 className="text-2xl my-4">Activities</h2>
      <div className="grid grid-cols-2 gap-2">
        <a
          href="https://github.com/aviator-co/celerymon"
          target="_blank"
          rel="noreferrer"
        >
          <Card
            title="celerymon"
            subtitle="Celery task queue monitoring tool"
          />
        </a>
        <a
          href="https://github.com/aviator-co/niche-git"
          target="_blank"
          rel="noreferrer"
        >
          <Card title="niche-git" subtitle="niche Git utility" />
        </a>
        <a
          href="https://github.com/google/hprof-parser"
          target="_blank"
          rel="noreferrer"
        >
          <Card title="hprof-parser" subtitle="JVM heap dump parser" />
        </a>
        <a
          href="https://github.com/google/goblet"
          target="_blank"
          rel="noreferrer"
        >
          <Card title="goblet" subtitle="Git caching proxy" />
        </a>
        <a
          href="https://github.com/google/gitprotocolio"
          target="_blank"
          rel="noreferrer"
        >
          <Card
            title="gitprotocolio"
            subtitle="Git protocol parser written in Go"
          />
        </a>
        <a
          href="https://github.com/google/ijaas"
          target="_blank"
          rel="noreferrer"
        >
          <Card title="ijaas" subtitle="IntelliJ as a Service" />
        </a>
      </div>
    </div>
  );
}

function Education() {
  return (
    <div>
      <h2 className="text-2xl my-4">Education</h2>
      <div className="space-y-2">
        <Card
          title="MS Computer Science Tokyo Institute of Technology"
          subtitle="March 2014. Takuo Watanabe Lab."
        >
          <div className="divide-y">
            <div className="flex gap-x-8 py-2">
              <div className="shrink-0 w-[10ch] text-right">Concentration</div>
              <p>Model checking and Fault tolerance</p>
            </div>
            <div className="flex gap-x-8 py-2">
              <div className="shrink-0 w-[10ch] text-right">Thesis</div>
              <div className="space-y-4">
                <p>
                  Full-Automatic Exhaustive Fault-Injection on Software Models
                  of Message-Passing Systems
                </p>

                <p>
                  Fault tolerance of distributed systems can be effectively
                  verified by model checking and fault injection, but its
                  process is highly error-prone. I proposed a way to solve this
                  problem by adding a language support to modeling languages,
                  which is a common approach in programming languages.
                </p>
              </div>
            </div>
          </div>
        </Card>
        <Card
          title="BS Computer Science Tokyo Institute of Technology"
          subtitle="March 2012. Takuo Watanabe Lab."
        >
          <div className="divide-y">
            <div className="flex gap-x-8 py-2">
              <div className="shrink-0 w-[10ch] text-right">Concentration</div>
              <p>Context-oriented programming</p>
            </div>

            <div className="flex gap-x-8 py-2">
              <div className="shrink-0 w-[10ch] text-right">Thesis</div>
              <div className="space-y-4">
                <p>
                  An Implementation Method of Context-Oriented Programming in
                  Objective-C
                </p>

                <p>
                  Context-oriented programming is a programming method that
                  enables us to define behaviors that depend on the
                  program&apos;s execution context. I proposed an implementation
                  method of Context-oriented programming in Objective-C.
                </p>
              </div>
            </div>
          </div>
        </Card>
      </div>
    </div>
  );
}

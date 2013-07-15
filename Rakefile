desc "Run jekyll"
task :jekyll do
  sh 'bundle exec jekyll build'
end

desc "Serve jekyll"
task :jekyll_serve do
  sh 'bundle exec jekyll serve --watch'
end

desc "Run compass"
task :compass do
  sh 'bundle exec compass compile -c _compass.rb --force'
end

desc "Serve compass"
task :compass_serve do
  sh 'bundle exec compass watch -c _compass.rb'
end

desc "Push to the server"
task :deploy do
  # commit = `git rev-list --max-count=1 origin/master`.strip
  # raise "rebase failed" unless $?
  # `git update-ref refs/heads/master #{commit}`
  # raise "rebase failed" unless $?
end

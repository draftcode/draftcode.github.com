desc "Run jekyll"
task :jekyll => :compass do
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

desc "Copy to _deploy"
task :deploy => :jekyll do
  cd '_deploy' do
    sh 'rm -r *'
    sh 'cp -R ../_site/* ./'
  end
end

desc "Push to the server"
task :push => :deploy do
  sh 'git push origin source:source'
  cd '_deploy' do
    sh 'git add -A'
    sh 'git commit -m "Update documentation"'
    sh 'git push origin master:master'
  end
end

ignore 'stylesheets/normalize.css'

activate :deploy do |deploy|
  deploy.method = :git
  deploy.branch = 'master'
end

configure :build do
  activate :minify_css
end

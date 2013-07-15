require 'redcarpet'

class Redcarpet::Render::HTML
  def header(text, level)
    level += 1
    "<h#{level}>#{text}</h#{level}>"
  end

  def normal_text(text)
    if text[0] == "\n"
      if /\w/ === text[1]
        ' ' + text[1..-1]
      else
        text[1..-1]
      end
    else
      text
    end
  end
end

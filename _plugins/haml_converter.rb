require "haml"

module Jekyll
  class HamlConverter < Converter
    safe true
    priority :low

    def matches(ext)
      ext =~ /haml/i
    end

    def output_ext(ext)
      ".html"
    end

    def convert(content)
      ::Haml::Engine.new(content).render
    end
  end
end

module Jekyll
  class Layout
    def content
      if ext == '.haml' && @converted != true
        if @content =~ /\A(---\s*\n.*?\n?)^(---\s*$\n?)/m
          @content = $& + ::Haml::Engine.new($').render
        else
          @content = ::Haml::Engine.new(@content).render
        end
        @converted = true
      end
      @content
    end
  end
end

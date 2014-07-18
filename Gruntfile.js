module.exports = function(grunt) {
    grunt.initConfig({
        newpost: {
        }
    });

    grunt.registerTask('newpost', "Generates a new template task with the given filename", function(name) {
        filePath = './content/post/'
        if (arguments.length == 0) {
            filename = "untitled.md";
            title = '"untitled"';
        } else {
            filename = name + '.md';
            title = '"' + name + '"';
        }
        
        defaultContent  = '---\n'
        defaultContent += 'title: ' + title + '\n'
        defaultContent += 'brief: "Enter a 2-3 liner here"\n'
        defaultContent += 'date: ' + generateTodayString() + '\n'
        defaultContent += 'type: post\n'
        defaultContent += '---\n\n'
        defaultContent += "Write your post here"

        // Generate the file
        fs = require('fs')
        fs.writeFileSync(filePath + filename, defaultContent)
        grunt.log.writeln('[' + this.name + '] Template post created. vim ' + filePath + filename)
    });

    grunt.registerTask('newproject', "Generates a new template project with the given filename", function(name) {
        filePath = './content/project/'
        if (arguments.length == 0) {
            filename = "untitled.md";
            title = '"untitled"';
        } else {
            filename = name + '.md';
            title = '"' + name + '"';
        }

        defaultContent  = '---\n'
        defaultContent += 'title: ' + title + '\n'
        defaultContent += 'brief: "Enter a 2-3 liner here"\n'
        defaultContent += 'date: ' + generateTodayString() + '\n'
        defaultContent += 'type: project\n'
        defaultContent += 'thumbnail: "/images/logo-something.png"\n'
        defaultContent += '---\n\n'
        defaultContent += "Write about your project here"

        // Generate the file
        fs = require('fs')
        fs.writeFileSync(filePath + filename, defaultContent)
        grunt.log.writeln('[' + this.name + '] Template project created. vim ' + filePath + filename)
    });

    function generateTodayString() {
        // Generate a date that hugo can understand
        today = new Date()
        year = today.getFullYear()
        month = today.getMonth()
        day = today.getDay()

        strDate  = year + '-'
        if(month<10) strDate += '0'
        strDate += month + '-'
        
        if(day<10) strDate += '0'
        strDate += day

        return strDate
    }
}

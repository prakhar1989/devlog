module.exports = function(grunt) {
    grunt.initConfig({
        newpost: {
        },
        deploy: {
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
        defaultContent += 'draft: true\n'
        defaultContent += '---\n\n'
        defaultContent += "Write your post here"

        // Generate the file
        fs = require('fs')
        fs.writeFileSync(filePath + filename, defaultContent)
        grunt.log.writeln('Template post created. vim ' + filePath + filename)
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
        defaultContent += 'draft: true\n'
        defaultContent += '---\n\n'
        defaultContent += "Write about your project here"

        // Generate the file
        fs = require('fs')
        fs.writeFileSync(filePath + filename, defaultContent)
        grunt.log.writeln('Template project created. vim ' + filePath + filename)
    });

    grunt.registerTask('generate', "Removes existing static website and regenerates it", function() {
        var sh = require('execSync')
        code = sh.run('hugo')
        if(code==0) {
            grunt.log.writeln('Blog generated successfully')
        } else {
            grunt.log.writeln('There was an error running hugo (' + code + ')')
        }
    });

    grunt.registerTask('deploy', "Deploys the latest set of files to my host", function() {
        var sh = require('execSync')
        var user = 'utkarsh'
        var server = 'utkarshsinha.com'
        var globalLocation = 'srv/'
        var publicLocation = 'public_html'
        var code = sh.run('rsync ./public/ ' + user + '@' + server + ':/' + globalLocation + server + '/' + publicLocation + ' -r')

        if(code==0) {
            grunt.log.writeln('The deploy was successful')
        } else {
            grunt.log.writeln('There was an error running rsync (' + code + ')')
        }
    });

    function generateTodayString() {
        // Generate a date that hugo can understand
        today = new Date()
        year = today.getFullYear()
        month = today.getMonth()
        day = today.getDate()

        strDate  = year + '-'
        if(month<10) strDate += '0'
        strDate += month + '-'
        
        if(day<10) strDate += '0'
        strDate += day

        return strDate
    }
}
